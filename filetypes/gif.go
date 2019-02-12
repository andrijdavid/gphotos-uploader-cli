package filetypes

import (
	"fmt"

	"github.com/juju/errors"
)

// isSameGifs checks if two gifs (local and uploaded) are exactly the same
func isSameGifs(upGifPath, localGifPath string) bool {
	upHash, err := fileHash(upGifPath)
	if err != nil {
		return false
	}
	localHash, err := fileHash(localGifPath)
	if err != nil {
		return false
	}

	return upHash == localHash
}

// GifTypedMedia implements TypedMedia for GIF files
type GifTypedMedia struct{}

// IsCorrectlyUploaded checks that the gif that was uploaded is the same as the local one, before deleting the local one
func (gm *GifTypedMedia) IsCorrectlyUploaded(uploadedFileURL, localFilePath string) (bool, error) {
	if !IsGif(localFilePath) {
		return false, fmt.Errorf("%s is not a gif. Not deleting local file", localFilePath)
	}

	// compare uploaded image and local one
	if isSameGifs(uploadedFileURL, localFilePath) {
		return true, nil
	}

	return false, errors.Errorf("Not sure if gif was uploaded correctly. Not deleting local file. URL: %s", uploadedFileURL)
}
