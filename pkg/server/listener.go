package server

import (
	"context"
	"crypto/tls"
	"go.ketch.com/lib/orlop/v2/errors"
	"go.ketch.com/lib/orlop/v2/log"
	"net"
)

// NewListener returns a net.Listener for the given config
func NewListener(ctx context.Context, config Config) (net.Listener, error) {
	l := log.WithContext(ctx).WithField("addr", config.Addr())

	// Start listening
	listener, err := net.Listen("tcp", config.Addr())
	if err != nil {
		l.WithError(err).Error("failed to listen")

		return nil, errors.Wrapf(
			err,
			"failed to listen on %s",
			config.Addr(),
		)
	}

	// If TLS is not enabled, return the listener as it is
	if !config.TLS.GetEnabled() {
		l.WithField("tls", "disabled").Info("listening")
		return listener, nil
	}

	// Since TLS is enabled, get the tls.Config and return a crypto/tls listener
	cfg, err := NewTLSConfig(ctx, config.TLS)
	if err != nil {
		_ = listener.Close()

		return nil, errors.Wrap(
			err,
			"failed to get server TLS config",
		)
	}

	l.WithField("tls", "enabled").Info("listening")
	return tls.NewListener(listener, cfg), nil
}
