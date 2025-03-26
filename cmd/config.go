package main

import (
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tydanny/foodwheel/internal/app/foodwheel/cuisines/cuisineadapter/bun"
	"github.com/tydanny/foodwheel/internal/log"
)

type configuration struct {
	App appConfig

	Database bun.Config

	Log log.Config
}

// appConfig is the configuration for the application.
type appConfig struct {
	// HTTP server address
	HttpAddr string

	// GRPC server address
	GrpcAddr string

	// Storage is the storage backend of the application
	Storage string
}

func (c configuration) Validate() error {
	return nil
}

func configure(appViper *viper.Viper, flagset *pflag.FlagSet) {
	// Viper Settings
	appViper.AddConfigPath(".")
	appViper.SetConfigName("config")
	appViper.SetConfigType("yaml")

	// Env Var Settings
	appViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	appViper.AllowEmptyEnv(true)
	appViper.AutomaticEnv()

	// Global Configuration
	appViper.SetDefault("shutdownTimeout", time.Second*15)

	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		appViper.SetDefault("no_color", true)
	}

	// Log Configuration
	appViper.SetDefault("log.format", "json")
	appViper.SetDefault("log.level", "info")
	appViper.RegisterAlias("log.noColor", "no_color")

	// App Configuration
	flagset.String("http-addr", ":8000", "App HTTP server address")
	_ = appViper.BindPFlag("app.httpAddr", flagset.Lookup("http-addr"))
	appViper.SetDefault("app.httpAddr", ":8000")

	flagset.String("grpc-addr", ":8001", "App GRPC server address")
	_ = appViper.BindPFlag("app.grpcAddr", flagset.Lookup("grpc-addr"))
	appViper.SetDefault("app.grpcAddr", ":8001")

	appViper.SetDefault("app.storage", "inmemory")

	// Database Configuration
	_ = appViper.BindEnv("Database.host")
	appViper.SetDefault("Database.port", 5432)
	_ = appViper.BindEnv("Database.user")
	_ = appViper.BindEnv("Database.pass")
	_ = appViper.BindEnv("Database.name")
	_ = appViper.BindEnv("Database.params")
}
