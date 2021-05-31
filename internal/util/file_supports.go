package util

import (
	"io"

	"github.com/gabriel-vasile/mimetype"
)

var SupportedMediaMIMEs = []string{
	"image/jpeg",
	"image/png",
	"image/gif",
}

func IsMediaSupported(r io.Reader) bool {
	mtype, err := mimetype.DetectReader(r)
	if err != nil {
		return false
	}
	return mimetype.EqualsAny(mtype.String(), SupportedMediaMIMEs...)
}
