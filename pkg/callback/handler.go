package callback

import (
	"context"
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ketch-com/ketch-forwarder/pkg/requestcontext"
	"github.com/ketch-com/ketch-forwarder/pkg/server"
	"github.com/ketch-com/ketch-forwarder/pkg/types"
	"go.ketch.com/lib/orlop/v2/errors"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer r.Body.Close()

	request := new(types.Request)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		server.WriteError(r.Context(), w, errors.Invalid(err))
		return
	}

	ctx = requestcontext.WithStandardObject(ctx, request)

	if err := validation.ValidateWithContext(ctx, request); err != nil {
		server.WriteError(ctx, w, errors.Invalid(err))
		return
	}

	if request.Kind == types.AccessStatusEventKind {
		accessRequest := &types.AccessStatusEvent{
			ApiVersion: request.ApiVersion,
			Kind:       request.Kind,
			Metadata:   request.Metadata,
		}

		if err := json.Unmarshal(request.Event, &accessRequest.Event); err != nil {
			server.WriteError(ctx, w, errors.Invalid(err))
			return
		}

		if err := validation.ValidateWithContext(ctx, accessRequest); err != nil {
			server.WriteError(ctx, w, errors.Invalid(err))
			return
		}

		if err := h.HandleAccessStatusEvent(ctx, accessRequest); err != nil {
			server.WriteError(ctx, w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	} else if request.Kind == types.DeleteStatusEventKind {
		deleteRequest := &types.DeleteStatusEvent{
			ApiVersion: request.ApiVersion,
			Kind:       request.Kind,
			Metadata:   request.Metadata,
		}

		if err := json.Unmarshal(request.Event, &deleteRequest.Event); err != nil {
			server.WriteError(ctx, w, errors.Invalid(err))
			return
		}

		if err := validation.ValidateWithContext(ctx, deleteRequest); err != nil {
			server.WriteError(ctx, w, errors.Invalid(err))
			return
		}

		if err := h.HandleDeleteStatusEvent(ctx, deleteRequest); err != nil {
			server.WriteError(ctx, w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	server.WriteError(ctx, w, errors.Invalidf("invalid request kind '%s'", request.Kind))
}

func (h *Handler) HandleAccessStatusEvent(ctx context.Context, request *types.AccessStatusEvent) error {
	fmt.Println(request)
	return nil
}

func (h *Handler) HandleDeleteStatusEvent(ctx context.Context, request *types.DeleteStatusEvent) error {
	fmt.Println(request)
	return nil
}
