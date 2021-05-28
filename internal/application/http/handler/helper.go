package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"

	"bakalo.li/internal/util"
)

func fetchIDFromParam(r *http.Request) (uint32, error) {
	idParam := chi.URLParam(r, "id")
	if idParam != "" {
		id, err := util.StrToUint32(idParam)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, errors.New("parameter 'id' is empty")
}
