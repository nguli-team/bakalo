package response

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"

	"github.com/nguli-team/bakalo/internal/domain"
)

type TokenRequestedResponse struct {
	RequestID   string `json:"request_id"`
	RequestedAt int64  `json:"requested_at"`
}

func NewTokenRequestedResponse() *TokenRequestedResponse {
	return &TokenRequestedResponse{}
}

func (rd *TokenRequestedResponse) Render(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	rd.RequestedAt = time.Now().Unix()
	rd.RequestID = middleware.GetReqID(ctx)
	return nil
}

type TokenLoggedInResponse struct {
	*domain.VipToken

	RequestID  string `json:"request_id"`
	LoggedInAt int64  `json:"logged_in_at"`
}

func NewTokenLoggedInResponse(token *domain.VipToken) *TokenLoggedInResponse {
	return &TokenLoggedInResponse{
		VipToken: token,
	}
}

func (rd *TokenLoggedInResponse) Render(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	rd.LoggedInAt = time.Now().Unix()
	rd.RequestID = middleware.GetReqID(ctx)
	return nil
}
