package internal

import (
	"github.com/ThousandMilesFirstStep/imagine/internal/filters"
	"github.com/ThousandMilesFirstStep/imagine/internal/models"
)

type Filter func(image *models.Image, conf map[string]interface{}) error

var filtersStore map[string]Filter

func init() {
	filtersStore = map[string]Filter{
		"thumbnail":  filters.Thumbnail,
		"autorotate": filters.Autorotate,
		"strip":      filters.Strip,
		"watermark":  filters.Watermark,
		"export":     filters.Export,
	}
}

func RegisterFilter(name string, filter Filter) {
	filtersStore[name] = filter
}

func getFilter(name string) Filter {
	return filtersStore[name]
}
