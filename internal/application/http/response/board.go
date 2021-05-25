package response

import (
	"bakalo.li/internal/domain"
	"github.com/go-chi/render"
	"net/http"
)

type BoardResponse struct {
	*domain.Board
}

func NewBoardResponse(board *domain.Board) *BoardResponse {
	return &BoardResponse{
		Board: board,
	}
}

func (rd *BoardResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewBoardListResponse(boards []*domain.Board) []render.Renderer {
	var list []render.Renderer
	for _, board := range boards {
		list = append(list, NewBoardResponse(board))
	}
	return list
}
