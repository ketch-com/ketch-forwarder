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
	response.Kind = types.ErrorKind
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
