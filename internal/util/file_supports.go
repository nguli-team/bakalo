package util

import (
	"github.com/h2non/filetype"
)

var SupportedMediaMIMEs = []string{
	"image/jpeg",
	"image/png",
	"image/gif",
}

func IsMediaSupported(buf []byte) bool {
	for _, mime := range SupportedMediaMIMEs {
		if filetype.IsMIME(buf, mime) {
			return true
		}
	}
	return false
}
