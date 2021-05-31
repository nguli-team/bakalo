package response

import (
	"net"
	"net/http"

	"github.com/go-chi/render"

	"github.com/nguli-team/bakalo/internal/domain"
)

type PostResponse struct {
	*domain.Post
	IsYou bool `json:"is_you"`
}

func NewPostResponse(post *domain.Post) *PostResponse {
	return &PostResponse{
		Post: post,
	}
}

func (rd *PostResponse) Render(w http.ResponseWriter, r *http.Request) error {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// ignore if ip is not found
		return nil
	}
	rd.IsYou = rd.IPv4 == ip

	return nil
}

func NewPostListResponse(posts []domain.Post) []render.Renderer {
	var list []render.Renderer
	for i, _ := range posts {
		list = append(list, NewPostResponse(&posts[i]))
	}
	return list
}
