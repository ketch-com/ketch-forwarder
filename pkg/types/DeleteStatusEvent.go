package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DeleteStatusEvent struct {
	ApiVersion string              `json:"apiVersion,omitempty"`
	Kind       Kind                `json:"kind,omitempty"`
	Metadata   *Metadata           `json:"metadata,omitempty"`
	Event      *DeleteResponseBody `json:"event,omitempty"`
}

func (r *DeleteStatusEvent) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(DeleteStatusEventKind)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Event, validation.Required),
	)
}

func (r *DeleteStatusEvent) GetApiVersion() string {
	return r.ApiVersion
}

func (r *DeleteStatusEvent) GetKind() Kind {
	return r.Kind
}

func (r *DeleteStatusEvent) GetMetadata() *Metadata {
	return r.Metadata
}
