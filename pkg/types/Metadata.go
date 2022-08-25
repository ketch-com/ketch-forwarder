package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Metadata struct {
	UID    string `json:"uid,omitempty"`
	Tenant string `json:"tenant,omitempty"`
}

func (r *Metadata) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.UID, validation.Required, is.UUID),
		validation.Field(&r.Tenant, validation.Required),
	)
}
