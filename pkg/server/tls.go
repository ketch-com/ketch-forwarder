package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"go.ketch.com/lib/orlop/v2/errors"
	"go.ketch.com/lib/orlop/v2/log"
	"os"
)

// NewTLSConfig returns a tls.Config for the given Config
func NewTLSConfig(ctx context.Context, cfg TLSConfig) (*tls.Config, error) {
	l := log.WithContext(ctx)

	l.WithField("enabled", cfg.GetEnabled()).Trace("creating new server TLS config")

	config := &tls.Config{
		ClientAuth: cfg.ClientAuth,
		MinVersion: tls.VersionTLS12,
	}

	if !strSliceContains(config.NextProtos, "http/1.1") {
		// Enable HTTP/1.1
		config.NextProtos = append(config.NextProtos, "http/1.1")

		l.Trace("enabling http/1.1")
	}

	if !strSliceContains(config.NextProtos, "h2") {
		// Enable HTTP/2
		config.NextProtos = append([]string{"h2"}, config.NextProtos...)

		l.Trace("enabling h2")
	}

	if !cfg.GetEnabled() {
		l.Trace("tls is disabled")
		return config, nil
	} else {
		l.Trace("tls is enabled")
	}

	l.Trace("tls configured manually")

	l.WithField("file", cfg.CertFile).Trace("loading certificate")

	certPEMBlock, err := os.ReadFile(cfg.CertFile)
	if err != nil {
		return nil, errors.Configuration(
			errors.WithUserMessage(
				errors.Wrap(
					err,
					"tls: failed to load certificate",
				),
				"failed to load certificate",
			),
		)
	}

	l.Trace("certificate loaded")

	l.WithField("file", cfg.KeyFile).Trace("loading key")

	keyPEMBlock, err := os.ReadFile(cfg.KeyFile)
	if err != nil {
		return nil, errors.Configuration(
			errors.WithUserMessage(
				errors.Wrap(
					err,
					"tls: failed to load private key",
				),
				"failed to load private key",
			),
		)
	}

	l.Trace("key loaded")

	config.ClientCAs = x509.NewCertPool()

	if len(cfg.RootCAFile) > 0 {
		l.WithField("file", cfg.RootCAFile).Trace("loading rootca")

		rootcaPEMBlock, err := os.ReadFile(cfg.RootCAFile)
		if err != nil {
			return nil, errors.Configuration(
				errors.WithUserMessage(
					errors.Wrap(
						err,
						"tls: failed to load RootCA certificates",
					),
					"failed to load Root CA certificates",
				),
			)
		}

		l.Trace("rootca loaded")

		if !config.ClientCAs.AppendCertsFromPEM(rootcaPEMBlock) {
			return nil, errors.Configuration(
				errors.WithUserMessage(
					errors.Wrap(
						err,
						"tls: failed to append RootCA certificates",
					),
					"failed to append RootCA certificates",
				),
			)
		}
	}

	l.Trace("making x509 keypair")

	c, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, errors.WithUserMessage(
			errors.Wrap(
				err,
				"tls: failed creating key pair",
			),
			"failed to create Key Pair",
		)
	}

	l.Trace("adding x509 keypair to certificate list")

	config.Certificates = append(config.Certificates, c)

	return config, nil
}

func strSliceContains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}
