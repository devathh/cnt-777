package images

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
)

type imageConverter struct{}

type ImageConverter interface {
	Base64ToImage(base64Image string) (image.Image, error)
}

func NewConverter() ImageConverter {
	return &imageConverter{}
}

func (c *imageConverter) Base64ToImage(base64Image string) (image.Image, error) {
	imgData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64string: %w", err)
	}

	img, err := jpeg.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image from bytes: %w", err)
	}

	return img, nil
}
