package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"syscall"
	"time"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/oklog/run"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tydanny/foodwheel/internal/log"
	"google.golang.org/grpc"
)

// Provisioned by ldflags.
//
//nolint:gochecknoglobals
var (
	GitCommit = "NOCOMMIT"
	GoVersion = runtime.Version()
	BuildDate = ""
)

func initializeVersion() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		os.Exit(1)
	}

	modified := false

	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			GitCommit = setting.Value
		case "vcs.time":
			BuildDate = setting.Value
		case "vcs.modified":
			modified = true
		}
	}

	if modified {
		GitCommit += "+CHANGES"
	}
}

const (
	// appName is an identifier-like name used anywhere this app needs to be identified.
	//
	// It identifies the application itself, the actual instance needs to be identified via environment
	// and other details.
	appName = "foodwheel"

	// friendlyAppName is the visible name of the application.
	friendlyAppName = "The Food Wheel"
)

//nolint:funlen
func main() {
	initializeVersion()

	appViper := viper.New()
	flagset := pflag.NewFlagSet(appName, pflag.ExitOnError)
	configure(appViper, flagset)

	flagset.String("config", "", "config file (default is ./config.yaml)")
	flagset.Bool("version", false, "Show version information")

	_ = flagset.Parse(os.Args[1:])

	if v, _ := flagset.GetBool("version"); v {
		printVersion()

		os.Exit(0)
	}

	if c, _ := flagset.GetString("config"); c != "" {
		appViper.SetConfigFile(c)
	}

	err := appViper.ReadInConfig()

	configFileNotFound := errors.As(err, &viper.ConfigFileNotFoundError{})
	if !configFileNotFound {
		emperror.Panic(errors.Wrap(err, "failed to read configuration"))
	}

	// Initialize config
	var config configuration
	err = appViper.Unmarshal(&config)
	emperror.Panic(errors.Wrap(err, "failed to unmarshal configuration"))

	logger := log.NewLogger(config.Log)
	logger = logger.With(
		slog.String("revision", GitCommit),
		slog.String("goVersion", GoVersion),
		slog.String("buildDate", BuildDate),
	)

	log.SetGlobalLogger(logger)

	if configFileNotFound {
		logger.Warn("config file not found")
	}

	err = config.Validate()
	if err != nil {
		logger.Error("config validation failed", slog.Any("error", err))

		os.Exit(3)
	}

	logger.Info("Starting application")

	// ctx := context.Background()

	// Connect to the Database
	// logger.Info("Connecting to database")
	// bunDB, err := bun.NewBunDB(ctx, config.Database)
	// emperror.Panic(errors.Wrap(err, "failed to create BunDB"))

	var group run.Group

	// Set up app server
	{
		logger := logger.With(slog.String("server", appName))

		httpRouter := mux.NewRouter()

		httpServer := &http.Server{
			Handler:           httpRouter,
			ReadHeaderTimeout: 10 * time.Second,
			ErrorLog:          slog.NewLogLogger(logger.Handler(), slog.LevelInfo),
		}
		defer func() {
			err := httpServer.Close()
			emperror.Panic(errors.Wrap(err, "failed to close HTTP server"))
		}()

		// Create a new GRPC server and defer its stop.
		grpcServer := grpc.NewServer()
		defer grpcServer.Stop()

		logger.Info(
			"listening on address",
			slog.String("address", config.App.GrpcAddr),
		)

		grpcLn, err := net.Listen("tcp", config.App.GrpcAddr)
		emperror.Panic(err)

		group.Add(
			func() error { return grpcServer.Serve(grpcLn) },
			func(_ error) { grpcServer.GracefulStop() },
		)
	}

	group.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

	_ = group.Run()
}

func printVersion() {
	//nolint:forbidigo
	fmt.Printf(
		"%s version %s (%s) built on %s\n",
		friendlyAppName, GoVersion, GitCommit, BuildDate,
	)
}
