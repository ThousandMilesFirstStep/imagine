package internal

import (
	"fmt"
	"math"

	"github.com/davidbyttow/govips/v2/vips"

	"github.com/ThousandMilesFirstStep/imagine/internal/models"
)

type ConfigVips struct {
	Concurrency  int
	CollectStats bool
	CacheTrace   bool
	ReportLeaks  bool
}

type ConfigFilter struct {
	Filter  string
	Options map[string]interface{}
}

type ConfigFilters map[string][]ConfigFilter

type Config struct {
	Quality int
	Filters ConfigFilters
	Vips    ConfigVips
}

func Process(image *models.Image, filter string, config *Config) ([]byte, error) {
	defer image.Image.Close()

	if config.Quality > 0 {
		image.Export.Quality = config.Quality
	}

	filterSteps := config.Filters[filter]

	for _, filterStep := range filterSteps {
		filterFunc := getFilter(filterStep.Filter)
		if filterFunc == nil {
			return nil, fmt.Errorf("the filter \"%s\" does not exist", filterStep.Filter)
		}

		err := filterFunc(image, filterStep.Options)
		if err != nil {
			return nil, err
		}
	}

	var bytes []byte
	var err error
	switch image.Export.Format {
	case "jpeg":
		if image.Image.HasAlpha() {
			image.Image.Flatten(&vips.Color{R: 255, G: 255, B: 255})
		}

		bytes, _, err = image.Image.ExportJpeg(&vips.JpegExportParams{
			Quality:        image.Export.Quality,
			Interlace:      image.Export.Interlace,
			OptimizeCoding: true,
			StripMetadata:  image.Export.Strip,
		})
	case "png":
		bytes, _, err = image.Image.ExportPng(&vips.PngExportParams{
			Compression:   pngCompressionFromQuality(image.Export.Quality),
			Interlace:     image.Export.Interlace,
			StripMetadata: image.Export.Strip,
		})
	case "webp":
		bytes, _, err = image.Image.ExportWebp(&vips.WebpExportParams{
			Quality:         image.Export.Quality,
			Lossless:        image.Export.Lossless,
			ReductionEffort: image.Export.Effort,
			StripMetadata:   image.Export.Strip,
		})
	}

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func pngCompressionFromQuality(quality int) int {
	return int(math.Max(0, math.Ceil(9-float64(quality)*0.1)))
}
