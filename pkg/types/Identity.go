package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Identity struct {
	IdentitySpace  string `json:"identitySpace,omitempty"`
	IdentityFormat string `json:"identityFormat,omitempty"`
	IdentityValue  string `json:"identityValue,omitempty"`
}

func (r *Identity) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.IdentitySpace, validation.Required),
		validation.Field(&r.IdentityFormat, validation.When(len(r.IdentityFormat) > 0, validation.In(IdentityFormats...))),
		validation.Field(&r.IdentityValue, validation.Required),
	)
}
