package response

import (
	"bakalo.li/internal/domain"
	"github.com/go-chi/render"
	"net/http"
)

type ThreadResponse struct {
	*domain.Thread
}

func NewThreadResponse(thread *domain.Thread) *ThreadResponse {
	return &ThreadResponse{
		Thread: thread,
	}
}

func (rd *ThreadResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewThreadListResponse(threads []*domain.Thread) []render.Renderer {
	var list []render.Renderer
	for _, thread := range threads {
		list = append(list, NewThreadResponse(thread))
	}
	return list
}
