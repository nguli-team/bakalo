package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"bakalo.li/internal/application/http/response"
	"bakalo.li/internal/domain"
	"bakalo.li/internal/storage"
)

type BoardHandler struct {
	boardService domain.BoardService
}

func NewBoardHandler(boardService domain.BoardService) *BoardHandler {
	return &BoardHandler{
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

	if len(boards) == 0 {
		render.JSON(w, r, make([]interface{}, 0))
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

	id, err := fetchIDFromParam(r)
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	board, err := h.boardService.FindByID(ctx, id)
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
		switch err {
		case storage.ErrRecordNotFound:
			_ = render.Render(w, r, response.ErrNotFound())
			break
		default:
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
