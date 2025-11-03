package services

import (
	"fmt"
	"image"

	"github.com/cnt-777/internal/domain/valueobjects"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

type barcodeScanner struct{}

type BarcodeScanner interface {
	ScanImage(img image.Image) (*valueobjects.Barcode, error)
}

func NewBarcodeScanner() BarcodeScanner {
	return &barcodeScanner{}
}

func (s *barcodeScanner) ScanImage(img image.Image) (*valueobjects.Barcode, error) {
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return nil, fmt.Errorf("failed to create *gozxing.BinaryBitmap: %w", err)
	}

	reader := oned.NewCode128Reader()
	res, _ := reader.Decode(bmp, nil)

	return valueobjects.NewBarcode(res.GetText())
}
