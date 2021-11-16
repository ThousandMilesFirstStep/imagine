package imagine

import (
	"errors"
	"io"

	"github.com/BurntSushi/toml"
	"github.com/davidbyttow/govips/v2/vips"

	"github.com/ThousandMilesFirstStep/imagine/internal"
	"github.com/ThousandMilesFirstStep/imagine/internal/models"
)

var config *internal.Config

func Init(configFile string) error {
	if configFile == "" {
		return errors.New("no config file provided")
	}

	config = &internal.Config{}

	_, err := toml.DecodeFile(configFile, config)
	if err != nil {
		return err
	}

	// Start vips with given configuration
	vipConfig := &vips.Config{
		ConcurrencyLevel: config.Vips.Concurrency,
		CollectStats:     config.Vips.CollectStats,
		CacheTrace:       config.Vips.CacheTrace,
		ReportLeaks:      config.Vips.ReportLeaks,
	}
	vips.Startup(vipConfig)

	return nil
}

func Shutdown() {
	vips.Shutdown()
}

func RegisterFilter(name string, filter internal.Filter) {
	internal.RegisterFilter(name, filter)
}

func ProcessFile(path string, filter string) ([]byte, error) {
	image, err := vips.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	return process(image, filter, config)
}

func ProcessReader(r io.Reader, filter string) ([]byte, error) {
	image, err := vips.NewImageFromReader(r)
	if err != nil {
		return nil, err
	}

	return process(image, filter, config)
}

func ProcessBuffer(buffer []byte, filter string) ([]byte, error) {
	image, err := vips.NewImageFromBuffer(buffer)
	if err != nil {
		return nil, err
	}

	return process(image, filter, config)
}

func process(image *vips.ImageRef, filter string, config *internal.Config) ([]byte, error) {
	img := models.NewImage(image)

	return internal.Process(img, filter, config)
}
