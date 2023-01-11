package interceptors

import (
	"context"

	"github.com/beyazit/twirp-boilerplate/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

func NewInterceptorMakeSmallHats() twirp.Interceptor {
	return func(next twirp.Method) twirp.Method {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			method, _ := twirp.MethodName(ctx)
			if method == "MakeHat" {
				return next(ctx, &haberdasher.Size{Inches: 1})
			}
			return next(ctx, req)
		}
	}
}
