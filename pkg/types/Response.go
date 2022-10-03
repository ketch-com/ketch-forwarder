package types

import (
	"context"
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Response struct {
	ApiVersion string          `json:"apiVersion,omitempty"`
	Kind       Kind            `json:"kind,omitempty"`
	Metadata   *Metadata       `json:"metadata,omitempty"`
	Response   json.RawMessage `json:"response,omitempty"`
	Event      json.RawMessage `json:"event,omitempty"`
}

func (r *Response) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(Kinds...)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Response, validation.When(r.Event == nil, validation.Required)),
		validation.Field(&r.Event, validation.When(r.Response == nil, validation.Required)),
	)
}

func (r *Response) GetApiVersion() string {
	return r.ApiVersion
}

func (r *Response) GetKind() Kind {
	return r.Kind
}

func (r *Response) GetMetadata() *Metadata {
	return r.Metadata
}
