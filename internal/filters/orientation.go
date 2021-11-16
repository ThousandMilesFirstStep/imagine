package filters

import (
	"github.com/ThousandMilesFirstStep/imagine/internal/models"
)

func Autorotate(image *models.Image, conf map[string]interface{}) error {
	return image.Image.AutoRotate()
}
