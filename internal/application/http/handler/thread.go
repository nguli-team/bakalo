package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"github.com/nguli-team/bakalo/internal/application/http/middleware"
	"github.com/nguli-team/bakalo/internal/application/http/request/media"
	"github.com/nguli-team/bakalo/internal/application/http/response"
	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/storage"
	"github.com/nguli-team/bakalo/internal/util"
)

type ThreadHandler struct {
	threadService domain.ThreadService
}

func NewThreadHandler(threadService domain.ThreadService) *ThreadHandler {
	return &ThreadHandler{
		threadService: threadService,
	}
}

func (h ThreadHandler) ListThreads(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var threads []domain.Thread
	var err error

	if qBoardID := r.URL.Query().Get("board_id"); qBoardID != "" {
		boardID, err := util.StrToUint32(qBoardID)
		if err != nil {
			tIDInvalidErr := errors.New("query 'board_id' is invalid")
			_ = render.Render(w, r, response.ErrInvalidRequest(tIDInvalidErr))
		}

		threads, err = h.threadService.FindByBoardID(ctx, boardID)
		if err != nil {
			_ = render.Render(w, r, response.ErrInternal(err))
			return
		}
	} else {
		threads, err = h.threadService.FindAll(ctx)
		if err != nil {
			_ = render.Render(w, r, response.ErrInternal(err))
			return
		}
	}

	if len(threads) == 0 {
		render.JSON(w, r, make([]interface{}, 0))
		return
	}

	err = render.RenderList(w, r, response.NewThreadListResponse(threads))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}

func (h ThreadHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := FetchIDFromParam(r, "id")
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	thread, err := h.threadService.FindByID(ctx, id)
	if err != nil {
		switch err {
		case storage.ErrRecordNotFound:
			_ = render.Render(w, r, response.ErrNotFound())
			break
		default:
			_ = render.Render(w, r, response.ErrInternal(err))
		}
		return
	}

	err = render.Render(w, r, response.NewThreadResponse(thread))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}

func (h ThreadHandler) CreateThreadMultipart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := r.ParseMultipartForm(5 << 20) // max size: 5MB
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	// parse request body
	boardID, err := util.StrToUint32(r.PostFormValue("board_id"))
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}
	title := r.PostFormValue("title")
	if title == "" {
		err := errors.New("'title' is missing")
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
	ip := middleware.GetRequestIP(ctx)

	// handle media upload
	filename, err := media.HandleUpload(r, "media")
	if err != nil {
		switch err {
		case media.ErrFileNotSupported:
		case media.ErrFileInvalid:
			_ = render.Render(w, r, response.ErrInvalidRequest(err))
			break
		default:
			_ = render.Render(w, r, response.ErrInternal(err))
		}
		return
	}

	threadRequest := &domain.Thread{
		BoardID: boardID,
		Title:   title,
		OP: &domain.Post{
			Name:          opName,
			Text:          opText,
			MediaFileName: filename,
			IPv4:          ip,
		},
	}

	// save thread
	thread, err := h.threadService.Create(ctx, threadRequest)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}

	err = render.Render(w, r, response.NewThreadResponse(thread))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}
