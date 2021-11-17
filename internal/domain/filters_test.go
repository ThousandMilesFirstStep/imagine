package domain

import (
	"testing"
)

type TestImage struct{}

func (ti *TestImage) Close() {}

func (ti *TestImage) SetExportOptions(options *ExportOptions) {}

func (ti *TestImage) Export() ([]byte, error) {
	return []byte{}, nil
}

func (ti *TestImage) Thumbnail(width int, height int, inset bool) error {
	return nil
}

func (ti *TestImage) Watermark(image string, position Position, padding int) error {
	return nil
}

func (ti *TestImage) Strip() error {
	return nil
}

func (ti *TestImage) AutoRotate() error {
	return nil
}

func (ti *TestImage) Flatten(color Color) error {
	return nil
}

func TestFlattenRunWithoutError(t *testing.T) {
	image := &TestImage{}
	config := map[string]interface{}{"backgroundColor": "#ffffff"}

	err := flatten(image, config)

	if err != nil {
		t.Fail()
	}
}
