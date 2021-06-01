package request

import (
	"net/http"

	"github.com/nguli-team/bakalo/internal/application/http/helper"
)

type VipTokenRequest struct {
	Email string `json:"email"`
	IP    string `json:"ip,omitempty"`
}

func (rd *VipTokenRequest) Bind(r *http.Request) error {
	ctx := r.Context()
	ip := helper.GetRequestIP(ctx)

	rd.IP = ip
	return nil
}

type VipLoginRequest struct {
	Token string `json:"token"`
	PIN   int    `json:"pin"`
	IP    string `json:"-,omitempty"`
}

func (rd *VipLoginRequest) Bind(r *http.Request) error {
	ctx := r.Context()
	ip := helper.GetRequestIP(ctx)

	rd.IP = ip
	return nil
}
