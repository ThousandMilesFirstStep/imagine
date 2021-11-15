package internal

import (
	"github.com/ThousandMilesFirstStep/imagine/internal/filters"
	"github.com/davidbyttow/govips/v2/vips"
)

type Filter func(image *vips.ImageRef, conf map[string]interface{}) error

var filtersStore map[string]Filter

func init() {
	filtersStore = map[string]Filter{
		"thumbnail":  filters.Thumbnail,
		"autorotate": filters.Autorotate,
		"strip":      filters.Strip,
		"watermark":  filters.Watermark,
	}
}

func RegisterFilter(name string, filter Filter) {
	filtersStore[name] = filter
}

func getFilter(name string) Filter {
	return filtersStore[name]
}
