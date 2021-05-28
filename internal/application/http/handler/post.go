package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"bakalo.li/internal/application/http/response"
	"bakalo.li/internal/domain"
	"bakalo.li/internal/util"
)

type PostHandler struct {
	postService domain.PostService
}

func NewPostHandler(postService domain.PostService) PostHandler {
	return PostHandler{
		postService: postService,
	}
}

func (h PostHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var posts []domain.Post
	var err error

	if qThreadID := r.URL.Query().Get("thread_id"); qThreadID != "" {
		threadID, err := util.StrToUint32(qThreadID)
		if err != nil {
			tIDInvalidErr := errors.New("query 'thread_id' is invalid")
			_ = render.Render(w, r, response.ErrInvalidRequest(tIDInvalidErr))
		}

		posts, err = h.postService.FindByThreadID(ctx, threadID)
		if err != nil {
			_ = render.Render(w, r, response.ErrInternal(err))
			return
		}
	} else {
		posts, err = h.postService.FindAll(ctx)
		if err != nil {
			_ = render.Render(w, r, response.ErrInternal(err))
			return
		}
	}

	err = render.RenderList(w, r, response.NewPostListResponse(posts))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}
