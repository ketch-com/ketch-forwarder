package server

import (
	"fmt"
	"go.ketch.com/lib/orlop/v2/config"
	"go.uber.org/fx"
)

// Option returns a server configuration
func Option(name string) fx.Option {
	annotation := fmt.Sprintf(`name:"%s"`, name)

	return fx.Options(
		config.Option[Config](name, name),

		fx.Provide(
			fx.Annotate(
				NewListener,
				fx.ParamTags(``, annotation),
				fx.ResultTags(annotation),
			),
			fx.Annotate(
				NewServer,
				fx.ParamTags(``, annotation),
				fx.ResultTags(annotation),
			),
		),
	)
}
