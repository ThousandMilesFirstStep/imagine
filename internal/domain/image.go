package domain

import (
	"errors"
	"regexp"
	"strconv"
)

type Position string

const (
	BottomRight Position = "bottom-right"
	BottomLeft  Position = "bottom-left"
	TopRight    Position = "top-right"
	TopLeft     Position = "top-left"
	Center      Position = "center"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

func NewColorFromHex(hex string) (Color, error) {
	matched, err := regexp.MatchString("^#[0-9a-fA-F]{6}$", hex)
	if err != nil {
		return Color{}, err
	}

	if !matched {
		return Color{}, errors.New("invalid color value")
	}

	red, _ := strconv.ParseInt(hex[1:3], 16, 16)
	green, _ := strconv.ParseInt(hex[3:5], 16, 16)
	blue, _ := strconv.ParseInt(hex[5:7], 16, 16)

	return Color{R: uint8(red), G: uint8(green), B: uint8(blue)}, nil
}

type ExportOptions struct {
	Quality   int
	Format    string
	Strip     bool
	Interlace bool // Jpeg / PNG only
	Lossless  bool // WebP only
	Effort    int  // WebP only
}

type Image interface {
	AutoRotate() error
	Close()
	Export() ([]byte, error)
	Flatten(color Color) error
	SetExportOptions(options *ExportOptions)
	Strip() error
	Thumbnail(width int, height int, inset bool) error
	Watermark(image string, position Position, padding int) error
}
