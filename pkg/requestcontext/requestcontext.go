package requestcontext

import (
	"context"
	"github.com/ketch-com/ketch-forwarder/pkg/types"
)

type ContextKey string

var (
	tenantKey     = ContextKey("tenant")
	uidKey        = ContextKey("uid")
	apiVersionKey = ContextKey("apiVersion")
	kindKey       = ContextKey("kind")
)

func Tenant(ctx context.Context) string {
	v, _ := ctx.Value(tenantKey).(string)
	return v
}

func WithTenant(ctx context.Context, tenant string) context.Context {
	return context.WithValue(ctx, tenantKey, tenant)
}

func UID(ctx context.Context) string {
	v, _ := ctx.Value(uidKey).(string)
	return v
}

func WithUID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, uidKey, uid)
}

func ApiVersion(ctx context.Context) string {
	v, _ := ctx.Value(apiVersionKey).(string)
	return v
}

func WithApiVersion(ctx context.Context, apiVersion string) context.Context {
	return context.WithValue(ctx, apiVersionKey, apiVersion)
}

func Kind(ctx context.Context) types.Kind {
	v, _ := ctx.Value(kindKey).(types.Kind)
	return v
}

func WithKind(ctx context.Context, kind types.Kind) context.Context {
	return context.WithValue(ctx, kindKey, kind)
}

func Metadata(ctx context.Context) *types.Metadata {
	return &types.Metadata{
		UID:    UID(ctx),
		Tenant: Tenant(ctx),
	}
}

func WithMetadata(ctx context.Context, metadata *types.Metadata) context.Context {
	if metadata == nil {
		return ctx
	}

	return WithUID(WithTenant(ctx, metadata.Tenant), metadata.UID)
}

func StandardObject(ctx context.Context) types.StandardObject {
	return &types.Request{
		ApiVersion: ApiVersion(ctx),
		Kind:       Kind(ctx),
		Metadata:   Metadata(ctx),
	}
}

func WithStandardObject(ctx context.Context, obj types.StandardObject) context.Context {
	if obj == nil {
		return ctx
	}

	return WithApiVersion(WithKind(WithMetadata(ctx, obj.GetMetadata()), obj.GetKind()), obj.GetApiVersion())
}
