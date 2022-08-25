package server

import (
	"go.uber.org/fx"
)

var Module = fx.Module("server",
	Option("server"),

	fx.Invoke(
		fx.Annotate(
			Lifecycle,
			fx.ParamTags(``, `name:"server"`, `name:"server"`),
		),
	),
)
