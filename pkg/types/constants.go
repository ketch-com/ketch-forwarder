package types

const (
	ApiVersion = "dsr/v1"
)

var (
	IdentityFormats = []any{"raw", "md5", "sha1", "sha256", "sha512"}
	Kinds           = []any{
		AccessRequestKind,
		AccessResponseKind,
		AccessStatusEventKind,
		CorrectionRequestKind,
		CorrectionResponseKind,
		CorrectionStatusEventKind,
		DeleteRequestKind,
		DeleteResponseKind,
		DeleteStatusEventKind,
		RestrictProcessingRequestKind,
		RestrictProcessingResponseKind,
		RestrictProcessingStatusEventKind,
		ErrorKind,
	}
	RequestStatuses = []any{
		UnknownRequestStatus,
		PendingRequestStatus,
		InProgressRequestStatus,
		CompletedRequestStatus,
		CancelledRequestStatus,
		DeniedRequestStatus,
	}
	RequestStatusReasons = []any{
		UnknownRequestStatusReason,
		//NeedUserVerificationRequestStatusReason,
		SuspectedFraudRequestStatusReason,
		InsufficientVerificationRequestStatusReason,
		NoMatchRequestStatusReason,
		ClaimNotCoveredRequestStatusReason,
		OutsideJurisdictionRequestStatusReason,
		TooManyRequestsRequestStatusReason,
		OtherRequestStatusReason,
	}
)

type Kind string

const (
	AccessRequestKind                 Kind = "AccessRequest"
	AccessResponseKind                Kind = "AccessResponse"
	AccessStatusEventKind             Kind = "AccessStatusEvent"
	CorrectionRequestKind             Kind = "CorrectionRequest"
	CorrectionResponseKind            Kind = "CorrectionResponse"
	CorrectionStatusEventKind         Kind = "CorrectionStatusEvent"
	DeleteRequestKind                 Kind = "DeleteRequest"
	DeleteResponseKind                Kind = "DeleteResponse"
	DeleteStatusEventKind             Kind = "DeleteStatusEvent"
	RestrictProcessingRequestKind     Kind = "RestrictProcessingRequest"
	RestrictProcessingResponseKind    Kind = "RestrictProcessingResponse"
	RestrictProcessingStatusEventKind Kind = "RestrictProcessingStatusEvent"
	ErrorKind                         Kind = "Error"
)

type RequestStatus string

const (
	UnknownRequestStatus    RequestStatus = "unknown"
	PendingRequestStatus    RequestStatus = "pending"
	InProgressRequestStatus RequestStatus = "in_progress"
	CompletedRequestStatus  RequestStatus = "completed"
	CancelledRequestStatus  RequestStatus = "cancelled"
	DeniedRequestStatus     RequestStatus = "denied"
)

type RequestStatusReason string

const (
	UnknownRequestStatusReason                  RequestStatusReason = "unknown"
	SuspectedFraudRequestStatusReason           RequestStatusReason = "suspected_fraud"
	InsufficientVerificationRequestStatusReason RequestStatusReason = "insufficient_verification"
	NoMatchRequestStatusReason                  RequestStatusReason = "no_match"
	ClaimNotCoveredRequestStatusReason          RequestStatusReason = "claim_not_covered"
	OutsideJurisdictionRequestStatusReason      RequestStatusReason = "outside_jurisdiction"
	TooManyRequestsRequestStatusReason          RequestStatusReason = "too_many_requests"
	OtherRequestStatusReason                    RequestStatusReason = "other"
	//NeedUserVerificationRequestStatusReason     RequestStatusReason = "need_user_verification"
)
