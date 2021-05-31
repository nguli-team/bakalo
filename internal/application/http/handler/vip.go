package handler

import (
	"net/http"

	"github.com/nguli-team/bakalo/internal/domain"
)

type VIPHandler struct {
	tokenService domain.TokenService
}

func (h VIPHandler) CreateToken(w http.ResponseWriter, r *http.Request) {

}

func (h VIPHandler) Login(w http.ResponseWriter, r *http.Request) {

}
