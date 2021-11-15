package filters

import "github.com/davidbyttow/govips/v2/vips"

func Autorotate(image *vips.ImageRef, conf map[string]interface{}) error {
	return image.AutoRotate()
}
