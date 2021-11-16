package filters

import (
	"github.com/ThousandMilesFirstStep/imagine/internal/models"
	"github.com/davidbyttow/govips/v2/vips"
)

func Thumbnail(image *models.Image, conf map[string]interface{}) error {
	width := int(conf["width"].(int64))
	height := int(conf["height"].(int64))

	crop := vips.InterestingAll
	if conf["inset"] == true {
		crop = vips.InterestingNone
	}

	return image.Image.Thumbnail(width, height, crop)
}
