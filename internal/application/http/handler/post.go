package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"github.com/nguli-team/bakalo/internal/application/http/helper"
	"github.com/nguli-team/bakalo/internal/application/http/request/media"
	"github.com/nguli-team/bakalo/internal/application/http/response"
	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/util"
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

	if len(posts) == 0 {
		render.JSON(w, r, make([]interface{}, 0))
		return
	}

	err = render.RenderList(w, r, response.NewPostListResponse(posts))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}

func (h PostHandler) CreatePostMultipart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := r.ParseMultipartForm(5 << 20) // max size: 5MB
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	// parse request body
	threadID, err := util.StrToUint32(r.PostFormValue("thread_id"))
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}
	opText := r.PostFormValue("text")
	if opText == "" {
		err := errors.New("'text' is missing")
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}
	opName := r.PostFormValue("name")
	ip := helper.GetRequestIP(ctx)

	// handle media upload
	filename, err := media.HandleUpload(r, "media")
	if err != nil {
		switch err {
		case media.ErrFileInvalid:
			break
		case media.ErrFileNotSupported:
			_ = render.Render(w, r, response.ErrInvalidRequest(err))
			return
		default:
			_ = render.Render(w, r, response.ErrInternal(err))
			return
		}
	}

	postRequest := &domain.Post{
		ThreadID:      threadID,
		Name:          opName,
		Text:          opText,
		MediaFileName: filename,
		IPv4:          ip,
	}

	// save thread
	post, err := h.postService.Create(ctx, postRequest)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}

	err = render.Render(w, r, response.NewPostResponse(post))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}
