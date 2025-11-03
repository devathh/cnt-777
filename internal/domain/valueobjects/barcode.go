package valueobjects

type Barcode struct {
	value string
}

func NewBarcode(value string) (*Barcode, error) {
	return &Barcode{
		value: value,
	}, nil
}

func (b *Barcode) Value() string {
	return b.value
}
