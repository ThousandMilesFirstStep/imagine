package imagine

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ThousandMilesFirstStep/imagine/internal/adapters"
	"github.com/ThousandMilesFirstStep/imagine/internal/domain"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/go-yaml/yaml"
)

var config *domain.Config

func Init(configFile string) error {
	if configFile == "" {
		return errors.New("no config file provided")
	}

	config = &domain.Config{}

	loadConfigFile(configFile)

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

func RegisterFilter(name string, filter domain.Filter) {
	domain.RegisterFilter(name, filter)
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

func process(image *vips.ImageRef, filter string, config *domain.Config) ([]byte, error) {
	img := adapters.NewVipsImage(image)

	return domain.Process(img, filter, config)
}

func loadConfigFile(configFile string) error {
	// Check file existence
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return errors.New("the config file does not exist")
	}

	// TOML
	if strings.HasSuffix(configFile, ".toml") {
		_, err := toml.DecodeFile(configFile, config)
		if err != nil {
			return err
		}

		return nil
	}

	// YAML
	if strings.HasSuffix(configFile, ".yaml") || strings.HasSuffix(configFile, ".yml") {
		f, err := os.Open(configFile)
		if err != nil {
			return err
		}

		decoder := yaml.NewDecoder(f)
		decoder.Decode(config)

		f.Close()

		return nil
	}

	return errors.New("the config file format is not supported")
}
