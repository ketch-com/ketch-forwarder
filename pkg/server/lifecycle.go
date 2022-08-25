package server

import (
	"context"
	"go.ketch.com/lib/orlop/v2/log"
	"go.uber.org/fx"
	"net"
	"net/http"
)

// Lifecycle runs the HTTP Server
func Lifecycle(lifecycle fx.Lifecycle, listener net.Listener, server *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error { return start(listener, server) },
		OnStop:  func(ctx context.Context) error { return stop(ctx, server) },
	})
}

func start(listener net.Listener, server *http.Server) error {
	go func() {
		log.Info("serving")

		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("failed to start server")
		} else if err == http.ErrServerClosed {
			log.Info("server stopped")
		}
	}()

	return nil
}

func stop(ctx context.Context, server *http.Server) error {
	log.Info("stopping server")
	return server.Shutdown(ctx)
}
