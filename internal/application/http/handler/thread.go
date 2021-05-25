package handler

import (
	"bakalo.li/internal/application/http/response"
	"bakalo.li/internal/domain"
	"bakalo.li/internal/logger"
	"bakalo.li/internal/util"
	"fmt"
	"github.com/go-chi/render"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strconv"
)

type ThreadHandler struct {
	threadService domain.ThreadService
}

func NewThreadHandler(threadService domain.ThreadService) ThreadHandler {
	return ThreadHandler{
		threadService: threadService,
	}
}

func (h ThreadHandler) ListThreads(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	threads, err := h.threadService.FindAll(ctx)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}
	err = render.RenderList(w, r, response.NewThreadListResponse(threads))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}

func (h ThreadHandler) CreateThreadMultipart(w http.ResponseWriter, r *http.Request) {
	// FIXME: Ya Allah! what have I done
	ctx := r.Context()

	err := r.ParseMultipartForm(5 << 20) // max size: 5MB
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	// handle media upload
	media, mediaHeader, err := r.FormFile("media")
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}
	defer func(media multipart.File) {
		err = media.Close()
		if err != nil {
			logger.Log.Warn(err)
		}
	}(media)

	namePrefix := util.RandomAlphaNumString(6)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}
	filename := fmt.Sprintf("%s-%s", namePrefix, mediaHeader.Filename)
	dst, err := os.Create("media/" + filename)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}
	defer func(dst *os.File) {
		err = dst.Close()
		if err != nil {
			logger.Log.Warn(err)
		}
	}(dst)

	_, err = io.Copy(dst, media)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}

	// parse request body
	boardIDStr := r.PostFormValue("board_id")
	boardID, err := strconv.ParseUint(boardIDStr, 10, 32)
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}
	title := r.PostFormValue("title")
	opName := r.PostFormValue("name")
	opText := r.PostFormValue("text")

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}

	threadRequest := &domain.Thread{
		BoardID: uint32(boardID),
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
