package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AccessStatusEvent struct {
	ApiVersion string                 `json:"apiVersion,omitempty"`
	Kind       Kind                   `json:"kind,omitempty"`
	Metadata   *Metadata              `json:"metadata,omitempty"`
	Event      *AccessStatusEventBody `json:"event,omitempty"`
}

func (r *AccessStatusEvent) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(AccessStatusEventKind)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Event, validation.Required),
	)
}

func (r *AccessStatusEvent) GetApiVersion() string {
	return r.ApiVersion
}

func (r *AccessStatusEvent) GetKind() Kind {
	return r.Kind
}

func (r *AccessStatusEvent) GetMetadata() *Metadata {
	return r.Metadata
}

type AccessStatusEventBody struct {
	Status                      RequestStatus       `json:"status,omitempty"`
	Reason                      RequestStatusReason `json:"reason,omitempty"`
	ExpectedCompletionTimestamp int64               `json:"expectedCompletionTimestamp,omitempty"`
	Results                     []*Callback         `json:"results,omitempty"`
}

func (r *AccessStatusEventBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Status, validation.Required, validation.In(RequestStatuses...)),
		validation.Field(&r.Reason, validation.When(len(r.Reason) > 0, validation.In(RequestStatusReasons...))),
		validation.Field(&r.ExpectedCompletionTimestamp, validation.Required),
	)
}
