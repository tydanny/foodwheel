package bun

import (
	"github.com/mdobak/go-xerrors"
)

var (
	ErrMissingHost = xerrors.Message("database host is required")
	ErrMissingPort = xerrors.Message("database port is required")
	ErrMissingUser = xerrors.Message("database user is required")
	ErrMissingName = xerrors.Message("database name is required")
)
