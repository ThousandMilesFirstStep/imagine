package filters

import (
	"github.com/ThousandMilesFirstStep/imagine/internal/models"
)

func Export(image *models.Image, conf map[string]interface{}) error {
	quality := conf["quality"]
	if quality != nil {
		image.Export.Quality = int(quality.(int64))
	}

	format := conf["format"]
	if format != nil {
		image.Export.Format = format.(string)
	}

	strip := conf["strip"]
	if strip != nil {
		image.Export.Strip = strip.(bool)
	}

	interlace := conf["interlace"]
	if interlace != nil {
		image.Export.Interlace = interlace.(bool)
	}

	lossless := conf["lossless"]
	if lossless != nil {
		image.Export.Lossless = lossless.(bool)
	}

	effort := conf["effort"]
	if effort != nil {
		image.Export.Effort = effort.(int)
	}

	return nil
}
