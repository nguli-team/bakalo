package response

import (
	"net/http"

	"github.com/go-chi/render"

	"bakalo.li/internal/domain"
)

type PostResponse struct {
	*domain.Post
}

func NewPostResponse(post *domain.Post) *PostResponse {
	return &PostResponse{
		Post: post,
	}
}

func (rd *PostResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewPostListResponse(posts []domain.Post) []render.Renderer {
	var list []render.Renderer
	for _, post := range posts {
		list = append(list, NewPostResponse(&post))
	}
	return list
}
