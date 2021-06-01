package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"github.com/nguli-team/bakalo/internal/application/http/request"
	"github.com/nguli-team/bakalo/internal/application/http/response"
	"github.com/nguli-team/bakalo/internal/domain"
)

type VIPHandler struct {
	tokenService domain.TokenService
}

func NewTokenHandler(tokenService domain.TokenService) *VIPHandler {
	return &VIPHandler{
		tokenService: tokenService,
	}
}

func (h VIPHandler) RequestNewToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := &request.VipTokenRequest{}
	err := render.Bind(r, data)
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	_, err = h.tokenService.CreateNewToken(ctx, data.IP, data.Email)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}

	err = render.Render(w, r, response.NewTokenRequestedResponse())
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}

func (h VIPHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := &request.VipLoginRequest{}
	err := render.Bind(r, data)
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	tokenIsValid := h.tokenService.ValidateToken(ctx, data.Token, data.PIN)
	if !tokenIsValid {
		err := errors.New("token or pin is not valid")
		_ = render.Render(w, r, response.ErrUnauthorized(err))
		return
	}

	updatedToken, err := h.tokenService.UpdateTokenIP(ctx, data.IP, data.Token)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}

	err = render.Render(w, r, response.NewTokenLoggedInResponse(updatedToken))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}
