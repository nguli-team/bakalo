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
	ctx := r.Context()

	var board domain.Board
	var err error

	if boardID := chi.URLParam(r, "id"); boardID != "" {
		var id int64
		id, err = strconv.ParseInt(boardID, 10, 64)
		if err != nil {
			_ = render.Render(w, r, response.ErrInvalidRequest(err))
			return
		}
		board, err = h.service.FindByID(ctx, id)
	} else if boardShorthand := chi.URLParam(r, "shorthand"); boardShorthand != "" {
		board, err = h.service.FindByShorthand(ctx, boardShorthand)
	} else {
		_ = render.Render(w, r, response.ErrInvalidRequest(errors.New("id or shorthand is invalid")))
		return
	}

	// FIXME: This is catch all error, but not every error is indicating a record is not found in the repository
	if err != nil {
		_ = render.Render(w, r, response.ErrNotFound())
		return
	}

	_ = render.Render(w, r, response.NewBoardResponse(board))
}

func (h BoardHandler) GetByShorthand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var board domain.Board
	var err error

	if boardShorthand := chi.URLParam(r, "shorthand"); boardShorthand != "" {
		board, err = h.service.FindByShorthand(ctx, boardShorthand)
	}

	if errors.Is(err, storage.ErrRecordNotFound) {
		_ = render.Render(w, r, response.ErrNotFound())
		return
	}

	_ = render.Render(w, r, response.NewBoardResponse(board))
}
