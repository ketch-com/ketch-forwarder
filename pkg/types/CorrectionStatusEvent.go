package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CorrectionStatusEvent struct {
	ApiVersion string                  `json:"apiVersion,omitempty"`
	Kind       Kind                    `json:"kind,omitempty"`
	Metadata   *Metadata               `json:"metadata,omitempty"`
	Event      *CorrectionResponseBody `json:"event,omitempty"`
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
