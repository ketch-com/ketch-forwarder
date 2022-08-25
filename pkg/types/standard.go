package types

type StandardObject interface {
	GetApiVersion() string
	GetKind() Kind
	GetMetadata() *Metadata
}
