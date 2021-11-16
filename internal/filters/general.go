package filters

import (
	"github.com/ThousandMilesFirstStep/imagine/internal/models"
	"github.com/davidbyttow/govips/v2/vips"
)

const (
	bottomRight = "bottom-right"
	bottomLeft  = "bottom-left"
	topRight    = "top-right"
	topLeft     = "top-left"
	center      = "center"
)

func Watermark(image *models.Image, conf map[string]interface{}) error {
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
		y = image.Image.Height() - watermark.Height() - padding
	case topRight:
		x = image.Image.Width() - watermark.Width() - padding
		y = padding
	case bottomRight:
		x = image.Image.Width() - watermark.Width() - padding
		y = image.Image.Height() - watermark.Height() - padding
	case center:
		x = image.Image.Width()/2 - watermark.Width()/2
		y = image.Image.Height()/2 - watermark.Height()/2
	}

	return image.Image.Composite(watermark, vips.BlendModeAtop, x, y)
}

func Strip(image *models.Image, conf map[string]interface{}) error {
	return image.Image.RemoveMetadata()
}
