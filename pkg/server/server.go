package server

import (
	"context"
	"go.ketch.com/lib/orlop/v2/log"
	stdlog "log"
	"net"
	"net/http"
)

// NewServer returns a new http.Server
func NewServer(ctx context.Context, handler http.Handler) *http.Server {
	return &http.Server{
		Handler: handler,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
		ErrorLog: stdlog.New(log.New().WithField("system", "http").Writer(), "", 0),
	}
}
