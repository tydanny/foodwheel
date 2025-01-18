package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/tydanny/foodwheel/internal/config"
	"go.uber.org/zap"
)

type mode string

var _ pflag.Value = (*mode)(nil)

func (m mode) Set(s string) error {
	m = mode(s)
	return nil
}

func (m mode) Type() string {
	return "mode"
}

func (m mode) String() string {
	return string(m)
}

const (
	DEV  mode = "DEV"
	PROD mode = "PROD"
)

var (
// tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
// certFile   = flag.String("cert_file", "", "The TLS cert file")
// keyFile    = flag.String("key_file", "", "The TLS key file")
// jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of Cuisines")
// port = flag.Int("port", 50051, "The server port")
)

func main() {
	// Setup flags
	mode := DEV
	pflag.Var(&mode, "mode", "The server mode")

	pflag.Parse()

	// Setup logger
	var zapLog *zap.Logger
	var logErr error
	if mode == PROD {
		zapLog, logErr = zap.NewProduction()
	} else {
		zapLog, logErr = zap.NewDevelopment()
	}
	if logErr != nil {
		panic(fmt.Errorf("failed to initialize logger: %v", logErr))
	}

	zap.ReplaceGlobals(zapLog)

	//Initialize config
	config.InitializeConfig(zapLog)

	// // Setup listener
	// lis, lisErr := net.Listen("tcp", fmt.Sprintf(":%d", (*port)))
	// if lisErr != nil {
	// 	log.Error(lisErr, "failed to listen")
	// 	return
	// }
	// defer lis.Close()

	// var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		// *certFile = data.Path("x509/server_cert.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		// *keyFile = data.Path("x509/server_key.pem")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }

	// grpcServer := grpc.NewServer()
	// reflection.Register(grpcServer)
	// foodwheel.RegisterFoodwheelServer(grpcServer, fwServer.ExampleServer())
	// log.V(1).Info("server listening", "mode", mode, "port", *port)
	// if serveErr := grpcServer.Serve(lis); serveErr != nil {
	// 	log.Error(serveErr, "failed to serve")
	// }
}
