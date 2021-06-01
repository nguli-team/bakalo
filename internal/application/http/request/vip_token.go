package request

import (
	"errors"
	"net/http"

	"github.com/nguli-team/bakalo/internal/application/http/middleware"
)

type VipTokenRequest struct {
	Email string `json:"email"`
	IP    string `json:"ip,omitempty"`
}

func (rd *VipTokenRequest) Bind(r *http.Request) error {
	ctx := r.Context()
	ip := middleware.GetRequestIP(ctx)
	if ip == "" {
		return errors.New("request IP is invalid")
	}

	rd.IP = ip
	return nil
}

type VipLoginRequest struct {
	Token string `json:"token"`
	IP    string `json:"ip,omitempty"`
}

func (rd *VipLoginRequest) Bind(r *http.Request) error {
	ctx := r.Context()
	ip := middleware.GetRequestIP(ctx)
	if ip == "" {
		return errors.New("request IP is invalid")
	}

	rd.IP = ip
	return nil
}
