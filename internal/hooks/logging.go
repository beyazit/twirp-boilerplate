package hooks

import (
	"context"
	"time"

	"github.com/twitchtv/twirp"

	log "github.com/sirupsen/logrus"
)

var ctxKey = new(int)

func NewLoggingServerHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			ctx = context.WithValue(ctx, ctxKey, time.Now())
			return ctx, nil
		},
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			service, _ := twirp.ServiceName(ctx)
			method, _ := twirp.MethodName(ctx)
			log.WithFields(log.Fields{
				"service": service,
				"method":  method,
			}).Info("request received")
			return ctx, nil
		},
		ResponseSent: func(ctx context.Context) {
			service, _ := twirp.ServiceName(ctx)
			method, _ := twirp.MethodName(ctx)
			log.WithFields(log.Fields{
				"service": service,
				"method":  method,
				"time":    time.Since(ctx.Value(ctxKey).(time.Time)),
			}).Info("response sent")
		},
		Error: func(ctx context.Context, twerr twirp.Error) context.Context {
			log.WithFields(log.Fields{
				"code": string(twerr.Code()),
			}).Info("error")
			return ctx
		},
	}
}
