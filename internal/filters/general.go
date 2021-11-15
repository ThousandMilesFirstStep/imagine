package filters

import (
	"github.com/davidbyttow/govips/v2/vips"
)

const (
	bottomRight = "bottom-right"
	bottomLeft  = "bottom-left"
	topRight    = "top-right"
	topLeft     = "top-left"
	center      = "center"
)

func Watermark(image *vips.ImageRef, conf map[string]interface{}) error {
	watermark, err := vips.NewImageFromFile(conf["image"].(string))
	if err != nil {
		return err
	}
	defer watermark.Close()

	position := topRight
	if conf["position"] != nil {
		position = conf["position"].(string)
	}

	padding := 0
	if conf["padding"] != nil {
		padding = int(conf["padding"].(int64))
	}

	var x, y int
	switch position {
	case topLeft:
		x, y = padding, padding
	case bottomLeft:
		x = padding
		y = image.Height() - watermark.Height() - padding
	case topRight:
		x = image.Width() - watermark.Width() - padding
		y = padding
	case bottomRight:
		x = image.Width() - watermark.Width() - padding
		y = image.Height() - watermark.Height() - padding
	case center:
		x = image.Width()/2 - watermark.Width()/2
		y = image.Height()/2 - watermark.Height()/2
	}

	return image.Composite(watermark, vips.BlendModeAtop, x, y)
}

func Strip(image *vips.ImageRef, conf map[string]interface{}) error {
	return image.RemoveMetadata()
}
