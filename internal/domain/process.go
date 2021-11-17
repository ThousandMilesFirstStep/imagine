package domain

import (
	"fmt"
	"math"
)

func Process(image Image, filter string, config *Config) ([]byte, error) {
	defer image.Close()

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

	bytes, err := image.Export()

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func PngCompressionFromQuality(quality int) int {
	return int(math.Max(0, math.Ceil(9-float64(quality)*0.1)))
}
