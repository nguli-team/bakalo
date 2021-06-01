package response

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/nguli-team/bakalo/internal/application/http/helper"
	"github.com/nguli-team/bakalo/internal/domain"
)

type ThreadResponse struct {
	*domain.Thread

	PostResponses []*PostResponse `json:"posts,omitempty"`
}

func NewThreadResponse(thread *domain.Thread) *ThreadResponse {
	return &ThreadResponse{
		Thread: thread,
	}
}

func (rd *ThreadResponse) Render(w http.ResponseWriter, r *http.Request) error {
	if len(rd.Posts) != 0 {
		ctx := r.Context()
		ip := helper.GetRequestIP(ctx)
		for _, post := range rd.Posts {
			pResponse := NewPostResponse(&post)
			pResponse.IsYou = ip == post.IPv4
			rd.PostResponses = append(rd.PostResponses, pResponse)
		}
	}
	return nil
}

func NewThreadListResponse(threads []domain.Thread) []render.Renderer {
	var list []render.Renderer
	for i, _ := range threads {
		list = append(list, NewThreadResponse(&threads[i]))
	}
	return list
}
