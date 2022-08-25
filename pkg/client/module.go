package client

import "go.uber.org/fx"

var Module = fx.Module("client",
	fx.Provide(
		NewCallbackClientProvider,
		NewForwarderClientProvider,
	),
)
