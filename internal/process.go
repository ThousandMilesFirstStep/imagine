package internal

import (
	"fmt"
	"math"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
)

type VipsConfig struct {
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
	Vips    VipsConfig
}

func Process(image *vips.ImageRef, filter string, config *Config) ([]byte, error) {
	defer image.Close()

	filterSteps := config.Filters[filter]

	for _, filterStep := range filterSteps {
		fmt.Println(filterStep.Filter, filterStep.Options)

		if getFilter(filterStep.Filter) == nil {
			return nil, fmt.Errorf("the filter \"%s\" does not exist", filterStep.Filter)
		}

		err := getFilter(filterStep.Filter)(image, filterStep.Options)
		if err != nil {
			return nil, err
		}
	}

	ext := strings.TrimLeft(image.Format().FileExt(), ".")

	var bytes []byte
	var err error
	switch ext {
	case "jpeg":
	case "jpg":
		bytes, _, err = image.ExportJpeg(&vips.JpegExportParams{
			Quality:        config.Quality,
			Interlace:      true,
			OptimizeCoding: true,
		})
	case "png":
		bytes, _, err = image.ExportPng(&vips.PngExportParams{
			Compression: qualityToCompression(config.Quality),
			Interlace:   false,
		})
	}

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func qualityToCompression(quality int) int {
	return int(math.Max(0, math.Ceil(9-float64(quality)*0.1)))
}
