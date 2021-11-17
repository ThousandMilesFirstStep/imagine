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

func GetCoordinatesFromPosition(w1 int, h1 int, w2 int, h2 int, position Position, padding int) (x int, y int) {
	switch position {
	case TopLeft:
		x, y = padding, padding
	case BottomLeft:
		x = padding
		y = h1 - h2 - padding
	case TopRight:
		x = w1 - w2 - padding
		y = padding
	case BottomRight:
		x = w1 - w2 - padding
		y = h1 - h2 - padding
	case Center:
		x = w1/2 - w2/2
		y = h1/2 - h2/2
	}

	return x, y
}
