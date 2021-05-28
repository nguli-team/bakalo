package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func fetchIDFromParam(r *http.Request) (uint32, error) {
	idParam := chi.URLParam(r, "id")
	if idParam != "" {
		var id64 uint64
		id64, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			return 0, err
		}
		id := uint32(id64)
		return id, nil
	}
	return 0, errors.New("parameter 'id' is empty")
}
