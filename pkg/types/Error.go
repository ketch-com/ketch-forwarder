package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Error struct {
	ApiVersion string     `json:"apiVersion,omitempty"`
	Kind       Kind       `json:"kind,omitempty"`
	Metadata   *Metadata  `json:"metadata,omitempty"`
	Error      *ErrorBody `json:"error,omitempty"`
}

func (r *Error) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(ErrorKind)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Error, validation.Required),
	)
}

type ErrorBody struct {
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *ErrorBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Code, validation.Required),
		validation.Field(&r.Status, validation.Required),
		validation.Field(&r.Message, validation.Required),
	)
}
