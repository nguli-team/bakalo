package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/nguli-team/bakalo/internal/util"
)

func FetchIDFromParam(r *http.Request, key string) (uint32, error) {
	idParam := chi.URLParam(r, key)
	if idParam != "" {
		id, err := util.StrToUint32(idParam)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, errors.New("parameter 'id' is empty")
}
