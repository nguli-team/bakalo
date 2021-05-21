package handler

import (
	"bakalo.li/internal/application/http/response"
	"bakalo.li/internal/domain"
	"bakalo.li/internal/logger"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

type BoardHandler struct {
	service domain.BoardService
}

func NewBoardHandler(boardService domain.BoardService) BoardHandler {
	return BoardHandler{
		service: boardService,
	}
}

func (h BoardHandler) ListBoards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	boards, _ := h.service.FindAll(ctx)

	err := render.RenderList(w, r, response.NewBoardListResponse(boards))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}

func (h BoardHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	logger.Log.Debug(r.RemoteAddr)
	render.Render(w, r, response.ErrInternal(errors.New("aksjdh")))
}
