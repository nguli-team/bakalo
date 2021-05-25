package handler

import (
	"bakalo.li/internal/application/http/response"
	"bakalo.li/internal/domain"
	"bakalo.li/internal/storage"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type BoardHandler struct {
	boardService domain.BoardService
}

func NewBoardHandler(boardService domain.BoardService) BoardHandler {
	return BoardHandler{
		boardService: boardService,
	}
}

func (h BoardHandler) ListBoards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	boards, err := h.boardService.FindAll(ctx)
	if err != nil {
		_ = render.Render(w, r, response.ErrInternal(err))
		return
	}

	err = render.RenderList(w, r, response.NewBoardListResponse(boards))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}

func (h BoardHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var board *domain.Board
	var err error

	if boardID := chi.URLParam(r, "id"); boardID != "" {
		var id64 uint64
		id64, err = strconv.ParseUint(boardID, 10, 32)
		if err != nil {
			_ = render.Render(w, r, response.ErrInvalidRequest(err))
			return
		}
		id := uint32(id64)
		board, err = h.boardService.FindByID(ctx, id)
	}

	if err != nil {
		if errors.Is(err, storage.ErrRecordNotFound) {
			_ = render.Render(w, r, response.ErrNotFound())
		} else {
			_ = render.Render(w, r, response.ErrInternal(err))
		}
		return
	}

	err = render.Render(w, r, response.NewBoardResponse(board))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}

func (h BoardHandler) GetByShorthand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var board *domain.Board
	var err error

	if boardShorthand := chi.URLParam(r, "shorthand"); boardShorthand != "" {
		board, err = h.boardService.FindByShorthand(ctx, boardShorthand)
	}

	if err != nil {
		if errors.Is(err, storage.ErrRecordNotFound) {
			_ = render.Render(w, r, response.ErrNotFound())
		} else {
			_ = render.Render(w, r, response.ErrInternal(err))
		}
		return
	}

	err = render.Render(w, r, response.NewBoardResponse(board))
	if err != nil {
		_ = render.Render(w, r, response.ErrRender(err))
		return
	}
}
