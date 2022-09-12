package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RestrictProcessingResponse struct {
	ApiVersion string                          `json:"apiVersion,omitempty"`
	Kind       Kind                            `json:"kind,omitempty"`
	Metadata   *Metadata                       `json:"metadata,omitempty"`
	Response   *RestrictProcessingResponseBody `json:"response"`
}

func (r *RestrictProcessingResponse) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(RestrictProcessingResponseKind)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Response, validation.Required),
	)
}

func (r *RestrictProcessingResponse) GetApiVersion() string {
	return r.ApiVersion
}

func (r *RestrictProcessingResponse) GetKind() Kind {
	return r.Kind
}

func (r *RestrictProcessingResponse) GetMetadata() *Metadata {
	return r.Metadata
}

type RestrictProcessingResponseBody struct {
	Status                      RequestStatus       `json:"status,omitempty"`
	Reason                      RequestStatusReason `json:"reason,omitempty"`
	ExpectedCompletionTimestamp int64               `json:"expectedCompletionTimestamp,omitempty"`
	RedirectURL                 string              `json:"redirectUrl,omitempty"`
	Results                     []*Callback         `json:"results,omitempty"`
}

func (r *RestrictProcessingResponseBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Status, validation.Required, validation.In(RequestStatuses...)),
		validation.Field(&r.Reason, validation.When(len(r.Reason) > 0, validation.In(RequestStatusReasons...))),
		validation.Field(&r.ExpectedCompletionTimestamp, validation.Required),
	)
}
