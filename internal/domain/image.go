package domain

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
