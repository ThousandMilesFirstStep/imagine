package adapters

import (
	"github.com/ThousandMilesFirstStep/imagine/internal/domain"
	"github.com/davidbyttow/govips/v2/vips"
)

type vipsImage struct {
	image         *vips.ImageRef
	exportOptions *domain.ExportOptions
}

func NewVipsImage(image *vips.ImageRef) *vipsImage {
	exportOptions := &domain.ExportOptions{
		Quality:   70,
		Interlace: true,
		Lossless:  false,
		Effort:    4,
		Strip:     true,
		Format:    vips.ImageTypes[image.Format()],
	}

	return &vipsImage{image, exportOptions}
}

func (vi *vipsImage) Close() {
	vi.image.Close()
}

func (vi *vipsImage) SetExportOptions(options *domain.ExportOptions) {
	vi.exportOptions = options
}

func (vi *vipsImage) Export() ([]byte, error) {
	var bytes []byte
	var err error

	switch vi.exportOptions.Format {
	case "jpeg":
		if vi.image.HasAlpha() {
			vi.image.Flatten(&vips.Color{R: 255, G: 255, B: 255})
		}

		bytes, _, err = vi.image.ExportJpeg(&vips.JpegExportParams{
			Quality:        vi.exportOptions.Quality,
			Interlace:      vi.exportOptions.Interlace,
			OptimizeCoding: true,
			StripMetadata:  vi.exportOptions.Strip,
		})
	case "png":
		bytes, _, err = vi.image.ExportPng(&vips.PngExportParams{
			Compression:   domain.PngCompressionFromQuality(vi.exportOptions.Quality),
			Interlace:     vi.exportOptions.Interlace,
			StripMetadata: vi.exportOptions.Strip,
		})
	case "webp":
		bytes, _, err = vi.image.ExportWebp(&vips.WebpExportParams{
			Quality:         vi.exportOptions.Quality,
			Lossless:        vi.exportOptions.Lossless,
			ReductionEffort: vi.exportOptions.Effort,
			StripMetadata:   vi.exportOptions.Strip,
		})
	}

	return bytes, err
}

func (vi *vipsImage) Thumbnail(width int, height int, inset bool) error {
	crop := vips.InterestingAll
	if inset == true {
		crop = vips.InterestingNone
	}

	return vi.image.Thumbnail(width, height, crop)
}

func (vi *vipsImage) Watermark(image string, position domain.Position, padding int) error {
	watermark, err := vips.NewImageFromFile(image)
	if err != nil {
		return err
	}
	defer watermark.Close()

	var x, y int
	switch position {
	case domain.TopLeft:
		x, y = padding, padding
	case domain.BottomLeft:
		x = padding
		y = vi.image.Height() - watermark.Height() - padding
	case domain.TopRight:
		x = vi.image.Width() - watermark.Width() - padding
		y = padding
	case domain.BottomRight:
		x = vi.image.Width() - watermark.Width() - padding
		y = vi.image.Height() - watermark.Height() - padding
	case domain.Center:
		x = vi.image.Width()/2 - watermark.Width()/2
		y = vi.image.Height()/2 - watermark.Height()/2
	}

	return vi.image.Composite(watermark, vips.BlendModeAtop, x, y)
}

func (vi *vipsImage) Strip() error {
	return vi.image.RemoveMetadata()
}

func (vi *vipsImage) AutoRotate() error {
	return vi.image.AutoRotate()
}

func (vi *vipsImage) Flatten(color domain.Color) error {
	return vi.image.Flatten((*vips.Color)(&color))
}
