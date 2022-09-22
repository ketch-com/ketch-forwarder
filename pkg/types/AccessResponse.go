package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AccessResponse struct {
	ApiVersion string              `json:"apiVersion,omitempty"`
	Kind       Kind                `json:"kind,omitempty"`
	Metadata   *Metadata           `json:"metadata,omitempty"`
	Response   *AccessResponseBody `json:"response"`
}

func (r *AccessResponse) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(AccessResponseKind)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Response, validation.Required),
	)
}

func (r *AccessResponse) GetApiVersion() string {
	return r.ApiVersion
}

func (r *AccessResponse) GetKind() Kind {
	return r.Kind
}

func (r *AccessResponse) GetMetadata() *Metadata {
	return r.Metadata
}

type AccessResponseBody struct {
	Status                      RequestStatus       `json:"status,omitempty"`
	Reason                      RequestStatusReason `json:"reason,omitempty"`
	RequestID                   string              `json:"requestID,omitempty"`
	ExpectedCompletionTimestamp int64               `json:"expectedCompletionTimestamp,omitempty"`
	Results                     []*Callback         `json:"results,omitempty"`
	//RedirectURL                 string              `json:"redirectUrl,omitempty"`
}

func (r *AccessResponseBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Status, validation.Required, validation.In(RequestStatuses...)),
		validation.Field(&r.Reason, validation.When(len(r.Reason) > 0, validation.In(RequestStatusReasons...))),
		validation.Field(&r.ExpectedCompletionTimestamp, validation.Required),
	)
}
