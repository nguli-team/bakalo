package response

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/nguli-team/bakalo/internal/domain"
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

func NewBoardListResponse(boards []domain.Board) []render.Renderer {
	var list []render.Renderer
	for i, _ := range boards {
		list = append(list, NewBoardResponse(&boards[i]))
	}
	return list
}
