package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"

	"github.com/valentin2105/Nemo/application"
	"github.com/valentin2105/Nemo/global"
)

var (
	addr    = flag.String("addr", ":8080", "Address / Port to bind on")
	tlscert = flag.String("tlscert", "", "TLS Certificate for HTTPS server (optional)")
	tlskey  = flag.String("tlskey", "", "TLS Private KEY for HTTPS server (optional)")
)

func newConfig() (*viper.Viper, error) {
	c := viper.New()
	c.SetDefault("cookie_secret", "xxe7wmUJN1H-7Lfa")
	c.SetDefault("http_drain_interval", "1s")
	c.AutomaticEnv()

	return c, nil
}

func main() {
	flag.Parse()
	config, err := newConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	app, err := application.New(config)
	if err != nil {
		logrus.Fatal(err)
	}
	middle, err := app.MiddlewareStruct()
	if err != nil {
		logrus.Fatal(err)
	}
	serverAddress := *addr
	certFile := *tlscert
	keyFile := *tlskey
	drainIntervalString := config.Get("http_drain_interval").(string)

	drainInterval, err := time.ParseDuration(drainIntervalString)
	if err != nil {
		logrus.Fatal(err)
	}

	srv := &graceful.Server{
		Timeout: drainInterval,
		Server:  &http.Server{Addr: serverAddress, Handler: middle},
	}

	version := global.GetEnv("KUBERNETES_VERSION", "v1.9")

	if certFile != "" && keyFile != "" {
		logrus.Infoln("Nemo is Running on -> " + serverAddress + " (HTTPS) for Kubernetes " + version)
		err = srv.ListenAndServeTLS(certFile, keyFile)
	} else {
		logrus.Infoln("Nemo is Running on -> " + serverAddress + " (HTTP) for Kubernetes " + version)
		err = srv.ListenAndServe()
	}

	if err != nil {
		logrus.Fatal(err)
	}
}
