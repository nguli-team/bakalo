package util

import (
	"io"
	"mime/multipart"

	"github.com/gabriel-vasile/mimetype"
)

var SupportedMediaMIMEs = []string{
	"image/jpeg",
	"image/png",
	"image/gif",
}

func IsMediaSupported(r multipart.File) bool {
	mtype, err := mimetype.DetectReader(r)
	_, _ = r.Seek(0, io.SeekStart)
	if err != nil {
		return false
	}
	return mimetype.EqualsAny(mtype.String(), SupportedMediaMIMEs...)
}
