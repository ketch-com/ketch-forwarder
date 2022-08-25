package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type DataSubject struct {
	Email           string `json:"email,omitempty"`
	FirstName       string `json:"firstName,omitempty"`
	LastName        string `json:"lastName,omitempty"`
	AddressLine1    string `json:"addressLine1,omitempty"`
	AddressLine2    string `json:"addressLine2,omitempty"`
	City            string `json:"city,omitempty"`
	StateRegionCode string `json:"stateRegionCode,omitempty"`
	PostalCode      string `json:"postalCode,omitempty"`
	CountryCode     string `json:"countryCode,omitempty"`
	Description     string `json:"description,omitempty"`
}

func (r *DataSubject) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.FirstName, validation.Required),
		validation.Field(&r.LastName, validation.Required),
	)
}
