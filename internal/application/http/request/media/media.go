package media

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/nguli-team/bakalo/internal/logger"
	"github.com/nguli-team/bakalo/internal/util"
)

// ErrFileInvalid indicates the file cannot be read properly.
var ErrFileInvalid = errors.New("media file is invalid")

// ErrFileNotSupported indicates the file is valid but the MIME type is not supported.
var ErrFileNotSupported = errors.New("media file is not supported")

// Writer writes media file to persistence storage.
// Defaults to write to disk.
var Writer = writeToDisk

// HandleUpload handles media upload and saves the file to storage.
// Takes an *http.Request and form key string to fetch the media file from.
// Returns the saved file name if success.
func HandleUpload(r *http.Request, formKey string) (string, error) {
	// parse media file
	mediaFile, mediaHandler, err := r.FormFile(formKey)
	if err != nil {
		return "", ErrFileInvalid
	}
	defer func(media multipart.File) {
		err = media.Close()
		if err != nil {
			logger.Log().Error(err)
		}
	}(mediaFile)

	// check file support
	supported := util.IsMediaSupported(mediaFile)
	if !supported {
		return "", ErrFileNotSupported
	}

	// set file name
	const prefixLength = 6
	namePrefix := util.RandomAlphaNumString(prefixLength) // generate random alphanumeric prefix
	filename := fmt.Sprintf("%s-%s", namePrefix, mediaHandler.Filename)

	// write file
	err = Writer(mediaFile, filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func writeToDisk(mediaFile io.Reader, filename string) error {
	// create write destination
	dst, err := os.Create("media/" + filename)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err = dst.Close()
		if err != nil {
			logger.Log().Error(err)
		}
	}(dst)

	// write file to destination
	_, err = io.Copy(dst, mediaFile)
	if err != nil {
		return err
	}

	return nil
}
