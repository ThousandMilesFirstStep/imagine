package models

import "github.com/davidbyttow/govips/v2/vips"

type ExportOptions struct {
	Quality   int
	Format    string
	Strip     bool
	Interlace bool // Jpeg / PNG only
	Lossless  bool // WebP only
	Effort    int  // WebP only
}

type Image struct {
	Image  *vips.ImageRef
	Export *ExportOptions
}

func NewImage(image *vips.ImageRef) *Image {
	ext := vips.ImageTypes[image.Format()]

	var export *ExportOptions
	switch ext {
	case "jpeg":
		export = defaultJpegExportOptions()
	case "png":
		export = defaultPngExportOptions()
	case "webp":
		export = defaultWebpExportOptions()
	default:
		export = defaultExportOptions()
	}

	export.Format = ext

	return &Image{
		Image:  image,
		Export: export,
	}
}

func defaultJpegExportOptions() *ExportOptions {
	return &ExportOptions{
		Quality:   70,
		Interlace: true,
		Strip:     true,
	}
}

func defaultPngExportOptions() *ExportOptions {
	return &ExportOptions{
		Quality:   70,
		Interlace: false,
		Strip:     true,
	}
}

func defaultWebpExportOptions() *ExportOptions {
	return &ExportOptions{
		Quality:  70,
		Lossless: false,
		Effort:   4,
		Strip:    true,
	}
}

func defaultExportOptions() *ExportOptions {
	return &ExportOptions{
		Quality:   70,
		Interlace: true,
		Lossless:  false,
		Effort:    4,
		Strip:     true,
	}
}
