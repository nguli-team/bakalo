package request

import (
	"bakalo.li/internal/domain"
	"errors"
	"net"
	"net/http"
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
