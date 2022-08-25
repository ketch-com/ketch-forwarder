package server

import (
	"context"
	"encoding/json"
	"github.com/ketch-com/ketch-forwarder/pkg/requestcontext"
	"github.com/ketch-com/ketch-forwarder/pkg/types"
	"go.ketch.com/lib/orlop/v2/errors"
	"go.ketch.com/lib/orlop/v2/log"
	"net/http"
)

func WriteError(ctx context.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(errors.StatusCode(err))
	response := new(types.Error)
	response.ApiVersion = types.ApiVersion
	response.Kind = MapErrorKind(ctx)
	response.Metadata = requestcontext.Metadata(ctx)
	response.Error = &types.ErrorBody{
		Code:    errors.StatusCode(err),
		Status:  errors.Code(err),
		Message: errors.UserMessage(err),
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.WithError(err).Error("failed to write error response")
	}
}

func MapErrorKind(ctx context.Context) types.Kind {
	switch requestcontext.Kind(ctx) {
	case types.AccessRequestKind, types.AccessResponseKind, types.AccessStatusEventKind, types.AccessErrorKind:
		return types.AccessErrorKind

	case types.DeleteRequestKind, types.DeleteResponseKind, types.DeleteStatusEventKind, types.DeleteErrorKind:
		return types.DeleteErrorKind

	default:
		return types.ErrorKind
	}
}
