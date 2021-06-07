package middleware

import (
	"context"
	"net/http"

	"github.com/nguli-team/bakalo/internal/application/http/helper"
	"github.com/nguli-team/bakalo/internal/domain"
)

func TokenRequest(tokenService domain.TokenService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			ip := helper.GetRequestIP(ctx)
			if ip == "" {
				next.ServeHTTP(w, r)
				return
			}

			token, err := tokenService.FindByIP(ctx, ip)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx = context.WithValue(ctx, domain.AdminContextKey, token.IsAdmin)
			ctx = context.WithValue(ctx, domain.VipContextKey, token.IsValid)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
