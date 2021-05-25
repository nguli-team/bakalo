package request

import (
	"errors"
	"net"
	"net/http"
)

type PostCreateRequest struct {
	Name string `json:"name"`
	Text string `json:"text"`
	IPv4 string `json:"-"`
}

func (rd PostCreateRequest) Bind(r *http.Request) error {
	if rd.Text != "" {
		return errors.New("post text field is missing")
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return errors.New("request is sent from invalid IP address")
	}
	rd.IPv4 = ip

	return nil
}
