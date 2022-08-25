//go:build !unit && !integration && smoke

package smoketest

import (
	"crypto/tls"
	"github.com/ketch-com/ketch-forwarder/pkg/client"
	"go.ketch.com/lib/orlop/v2/config"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Module("smoketest",
	fx.Supply(&http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		}},
	}),
	config.Option[Config](),
	client.Module,
)
