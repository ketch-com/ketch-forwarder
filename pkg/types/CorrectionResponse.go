package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CorrectionResponse struct {
	ApiVersion string                  `json:"apiVersion,omitempty"`
	Kind       Kind                    `json:"kind,omitempty"`
	Metadata   *Metadata               `json:"metadata,omitempty"`
	Response   *CorrectionResponseBody `json:"response,omitempty"`
}

func (r *CorrectionResponse) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(CorrectionResponseKind)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Response, validation.Required),
	)
}

func (r *CorrectionResponse) GetApiVersion() string {
	return r.ApiVersion
}

func (r *CorrectionResponse) GetKind() Kind {
	return r.Kind
}

func (r *CorrectionResponse) GetMetadata() *Metadata {
	return r.Metadata
}

type CorrectionResponseBody struct {
	Status                      RequestStatus       `json:"status,omitempty"`
	Reason                      RequestStatusReason `json:"reason,omitempty"`
	RequestID                   string              `json:"requestID,omitempty"`
	ExpectedCompletionTimestamp int64               `json:"expectedCompletionTimestamp,omitempty"`
	Claims                      map[string]any      `json:"claims,omitempty"`
	//RedirectURL                 string              `json:"redirectUrl,omitempty"`
}

func (r *CorrectionResponseBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Status, validation.Required, validation.In(RequestStatuses...)),
		validation.Field(&r.Reason, validation.When(len(r.Reason) > 0, validation.In(RequestStatusReasons...))),
		validation.Field(&r.ExpectedCompletionTimestamp, validation.Required),
	)
}
