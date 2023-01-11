package main

import (
	"net/http"

	"github.com/beyazit/twirp-boilerplate/internal/haberdasher"
	"github.com/beyazit/twirp-boilerplate/internal/hooks"
	"github.com/beyazit/twirp-boilerplate/internal/interceptors"
	rpc "github.com/beyazit/twirp-boilerplate/rpc/haberdasher"
	log "github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	svc := haberdasher.NewHaberdasherService()
	twirpHandler := rpc.NewHaberdasherServer(svc,
		twirp.WithServerInterceptors(interceptors.NewInterceptorMakeSmallHats()),
		twirp.WithServerHooks(hooks.NewLoggingServerHooks()))

	port := ":8080"
	log.WithField("port", port).Info("server started")
	http.ListenAndServe(port, twirpHandler)
}
