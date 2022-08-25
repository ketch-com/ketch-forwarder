package results

import (
	"github.com/go-chi/chi/v5"
	"github.com/ketch-com/ketch-forwarder/pkg/server"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Module("results",
	server.Module,

	fx.Provide(
		NewHandler,
		fx.Annotate(func(handler *Handler) http.Handler {
			mux := chi.NewMux()
			mux.Get("/results/{id}", handler.ServeHTTP)
			return mux
		}, fx.ResultTags(`name:"server"`)),
	),
)
