package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DeleteResponse struct {
	ApiVersion string              `json:"apiVersion,omitempty"`
	Kind       Kind                `json:"kind,omitempty"`
	Metadata   *Metadata           `json:"metadata,omitempty"`
	Response   *DeleteResponseBody `json:"response,omitempty"`
}

func (r *DeleteResponse) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(DeleteResponseKind)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Response, validation.Required),
	)
}

func (r *DeleteResponse) GetApiVersion() string {
	return r.ApiVersion
}

func (r *DeleteResponse) GetKind() Kind {
	return r.Kind
}

func (r *DeleteResponse) GetMetadata() *Metadata {
	return r.Metadata
}

type DeleteResponseBody struct {
	Status                      RequestStatus       `json:"status,omitempty"`
	Reason                      RequestStatusReason `json:"reason,omitempty"`
	RequestID                   string              `json:"requestID,omitempty"`
	ExpectedCompletionTimestamp int64               `json:"expectedCompletionTimestamp,omitempty"`
	//RedirectURL                 string              `json:"redirectUrl,omitempty"`
}

func (r *DeleteResponseBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Status, validation.Required, validation.In(RequestStatuses...)),
		validation.Field(&r.Reason, validation.When(len(r.Reason) > 0, validation.In(RequestStatusReasons...))),
		validation.Field(&r.ExpectedCompletionTimestamp, validation.Required),
	)
}
