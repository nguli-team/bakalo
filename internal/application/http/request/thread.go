package request

import (
	"errors"
	"net"
	"net/http"

	"bakalo.li/internal/domain"
)

type ThreadCreateRequest struct {
	Title string      `json:"title"`
	OP    domain.Post `json:"op"`
}

func (rd ThreadCreateRequest) Bind(r *http.Request) error {
	if rd.Title == "" {
		return errors.New("missing required thread title")
	}

	// Bind OP request
	if rd.OP.Text != "" {
		return errors.New("post text field is missing")
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return errors.New("request is sent from invalid IP address")
	}
	rd.OP.IPv4 = ip

	return nil
}
