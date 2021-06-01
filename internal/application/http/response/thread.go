package response

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/nguli-team/bakalo/internal/application/http/helper"
	"github.com/nguli-team/bakalo/internal/domain"
)

type ThreadResponse struct {
	*domain.Thread

	OPWithIsYou   *PostResponse   `json:"op"`
	PostResponses []*PostResponse `json:"posts,omitempty"`
}

func NewThreadResponse(thread *domain.Thread) *ThreadResponse {
	return &ThreadResponse{
		Thread: thread,
	}
}

func (rd *ThreadResponse) Render(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	ip := helper.GetRequestIP(ctx)

	rd.OPWithIsYou = NewPostResponse(rd.OP)
	rd.OPWithIsYou.IsYou = rd.OP.IPv4 == ip

	if len(rd.Posts) != 0 {
		for i := range rd.Posts {
			pResponse := NewPostResponse(&rd.Posts[i])
			pResponse.IsYou = ip == pResponse.IPv4
			rd.PostResponses = append(rd.PostResponses, pResponse)
		}
	}
	return nil
}

func NewThreadListResponse(threads []domain.Thread) []render.Renderer {
	var list []render.Renderer
	for i := range threads {
		list = append(list, NewThreadResponse(&threads[i]))
	}
	return list
}
