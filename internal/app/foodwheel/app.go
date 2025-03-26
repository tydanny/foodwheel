package foodwheel

import (
	"database/sql"
	"log/slog"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitializeFoodWheel(
	httpRouter *mux.Router,
	grpcServer *grpc.Server,
	storage string,
	db *sql.DB,
	logger *slog.Logger,
	errorHandler ErrorHandler,
) {
	// var store
}
