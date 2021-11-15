package filters

import "github.com/davidbyttow/govips/v2/vips"

func Thumbnail(image *vips.ImageRef, conf map[string]interface{}) error {
	width := int(conf["width"].(int64))
	height := int(conf["height"].(int64))

	crop := vips.InterestingAll
	if conf["inset"] == true {
		crop = vips.InterestingNone
	}

	return image.Thumbnail(width, height, crop)
}
