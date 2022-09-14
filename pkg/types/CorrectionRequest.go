package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CorrectionRequest struct {
	ApiVersion string                 `json:"apiVersion,omitempty"`
	Kind       Kind                   `json:"kind,omitempty"`
	Metadata   *Metadata              `json:"metadata,omitempty"`
	Request    *CorrectionRequestBody `json:"request,omitempty"`
}

func (r *CorrectionRequest) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(CorrectionRequestKind)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Request, validation.Required),
	)
}

func (r *CorrectionRequest) GetApiVersion() string {
	return r.ApiVersion
}

func (r *CorrectionRequest) GetKind() Kind {
	return r.Kind
}

func (r *CorrectionRequest) GetMetadata() *Metadata {
	return r.Metadata
}

type CorrectionRequestBody struct {
	Controller         string            `json:"controller,omitempty"`
	Property           string            `json:"property,omitempty"`
	Environment        string            `json:"environment,omitempty"`
	Regulation         string            `json:"regulation,omitempty"`
	Jurisdiction       string            `json:"jurisdiction,omitempty"`
	Identities         []*Identity       `json:"identities,omitempty"`
	Callbacks          []*Callback       `json:"callbacks,omitempty"`
	Subject            *DataSubject      `json:"subject,omitempty"`
	Claims             map[string]string `json:"claims,omitempty"`
	SubmittedTimestamp int64             `json:"submittedTimestamp,omitempty"`
	DueTimestamp       int64             `json:"dueTimestamp,omitempty"`
}

func (r *CorrectionRequestBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Property, validation.Required),
		validation.Field(&r.Environment, validation.Required),
		validation.Field(&r.Regulation, validation.Required),
		validation.Field(&r.Jurisdiction, validation.Required),
		validation.Field(&r.Identities, validation.Required),
		validation.Field(&r.Callbacks, validation.Required),
		validation.Field(&r.Subject, validation.Required),
		validation.Field(&r.SubmittedTimestamp, validation.Required),
	)
}
