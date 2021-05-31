package middleware

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/render"

	"github.com/nguli-team/bakalo/internal/application/http/response"
	"github.com/nguli-team/bakalo/internal/logger"
)

var IPContextKey = "request-ip"

func RequestIP(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil || ip == "" {
			logger.Log().Error(err)
			_ = render.Render(w, r, response.ErrInternal(err))
			return
		}

		ctx = context.WithValue(ctx, IPContextKey, ip)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func GetRequestIP(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if ip, ok := ctx.Value(IPContextKey).(string); ok {
		return ip
	}
	return ""
}
