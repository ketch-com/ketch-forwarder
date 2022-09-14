package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CorrectionStatusEvent struct {
	ApiVersion string                     `json:"apiVersion,omitempty"`
	Kind       Kind                       `json:"kind,omitempty"`
	Metadata   *Metadata                  `json:"metadata,omitempty"`
	Event      *CorrectionStatusEventBody `json:"event,omitempty"`
}

func (r *CorrectionStatusEvent) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(CorrectionStatusEventKind)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Event, validation.Required),
	)
}

func (r *CorrectionStatusEvent) GetApiVersion() string {
	return r.ApiVersion
}

func (r *CorrectionStatusEvent) GetKind() Kind {
	return r.Kind
}

func (r *CorrectionStatusEvent) GetMetadata() *Metadata {
	return r.Metadata
}

type CorrectionStatusEventBody struct {
	Status                      RequestStatus       `json:"status,omitempty"`
	Reason                      RequestStatusReason `json:"reason,omitempty"`
	ExpectedCompletionTimestamp int64               `json:"expectedCompletionTimestamp,omitempty"`
}

func (r *CorrectionStatusEventBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Status, validation.Required, validation.In(RequestStatuses...)),
		validation.Field(&r.Reason, validation.When(len(r.Reason) > 0, validation.In(RequestStatusReasons...))),
		validation.Field(&r.ExpectedCompletionTimestamp, validation.Required),
	)
}
