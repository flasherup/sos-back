package main

import (
	"crypto/tls"
	"github.com/flasherup/sos-back/server"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"os"

)


func main() {
	staticFolder := "./src/www"
	https := ":443"
	httpAddr := ":8888"
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "apisvc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	hs := server.NewHTTPTSTransport(logger, staticFolder)

	errs := make(chan error)
	go func() {
		level.Info(logger).Log("transport", "HTTPS", "addr", https)
		mgr := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("flasherup.com"),
			Cache:      autocert.DirCache(https + "cert/"), // to store certs
		}

		cfg := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			},
			GetCertificate: mgr.GetCertificate,
		}

		endpoint := &http.Server{
			Addr:    https,
			Handler: hs,
			TLSConfig: cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		errs <- endpoint.ListenAndServeTLS("", "")
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTPTest", "addr", httpAddr)
		endpoint := &http.Server{
			Addr:    httpAddr,
			Handler: hs,
		}
		errs <- endpoint.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
}