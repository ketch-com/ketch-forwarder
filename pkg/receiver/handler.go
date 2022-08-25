package receiver

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
	"time"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		server.WriteError(r.Context(), w, errors.Invalidf("expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type")))
		return
	}

	if !server.CanAccept(r, "application/json") {
		server.WriteError(r.Context(), w, errors.Invalidf("expected Accept to include 'application/json', got '%s'", r.Header.Get("Accept")))
		return
	}

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

	if request.Kind == types.AccessRequestKind {
		accessRequest := &types.AccessRequest{
			ApiVersion: request.ApiVersion,
			Kind:       request.Kind,
			Metadata:   request.Metadata,
		}

		if err := json.Unmarshal(request.Request, &accessRequest.Request); err != nil {
			server.WriteError(ctx, w, errors.Invalid(err))
			return
		}

		if err := validation.ValidateWithContext(ctx, accessRequest); err != nil {
			server.WriteError(ctx, w, errors.Invalid(err))
			return
		}

		resp, err := h.HandleAccessRequest(ctx, accessRequest)
		if err != nil {
			server.WriteError(ctx, w, err)
			return
		}

		if err := validation.ValidateWithContext(ctx, resp); err != nil {
			server.WriteError(ctx, w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(resp); err != nil {
			server.WriteError(ctx, w, err)
			return
		}

		return
	} else if request.Kind == types.DeleteRequestKind {
		deleteRequest := &types.DeleteRequest{
			ApiVersion: request.ApiVersion,
			Kind:       request.Kind,
			Metadata:   request.Metadata,
		}

		if err := json.Unmarshal(request.Request, &deleteRequest.Request); err != nil {
			server.WriteError(ctx, w, errors.Invalid(err))
			return
		}

		if err := validation.ValidateWithContext(ctx, deleteRequest); err != nil {
			server.WriteError(ctx, w, errors.Invalid(err))
			return
		}

		resp, err := h.HandleDeleteRequest(ctx, deleteRequest)
		if err != nil {
			server.WriteError(ctx, w, err)
			return
		}

		if err := validation.ValidateWithContext(ctx, resp); err != nil {
			server.WriteError(ctx, w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(resp); err != nil {
			server.WriteError(ctx, w, err)
			return
		}

		return
	}

	server.WriteError(ctx, w, errors.Invalidf("invalid request kind '%s'", request.Kind))
}

func (h *Handler) HandleAccessRequest(ctx context.Context, request *types.AccessRequest) (*types.AccessResponse, error) {
	fmt.Println(request)
	resp := &types.AccessResponse{
		ApiVersion: types.ApiVersion,
		Kind:       types.AccessResponseKind,
		Metadata:   request.Metadata,
		Response: &types.AccessResponseBody{
			Status:                      types.PendingRequestStatus,
			Reason:                      types.NeedUserVerificationRequestStatusReason,
			ExpectedCompletionTimestamp: time.Now().Add(45 * 24 * time.Hour).Unix(),
			RedirectURL:                 "https://idverification/123",
		},
	}

	return resp, nil
}

func (h *Handler) HandleDeleteRequest(ctx context.Context, request *types.DeleteRequest) (*types.DeleteResponse, error) {
	fmt.Println(request)
	resp := &types.DeleteResponse{
		ApiVersion: types.ApiVersion,
		Kind:       types.DeleteResponseKind,
		Metadata:   request.Metadata,
		Response: &types.DeleteResponseBody{
			Status:                      types.PendingRequestStatus,
			Reason:                      types.NeedUserVerificationRequestStatusReason,
			ExpectedCompletionTimestamp: time.Now().Add(45 * 24 * time.Hour).Unix(),
			RedirectURL:                 "https://idverification/123",
		},
	}
	return resp, nil
}
