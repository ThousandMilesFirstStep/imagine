package domain

type Filter func(image Image, conf map[string]interface{}) error

var filtersStore map[string]Filter

func init() {
	filtersStore = map[string]Filter{
		"thumbnail":  thumbnail,
		"autorotate": autorotate,
		"strip":      strip,
		"watermark":  watermark,
		"export":     setExportOptions,
		"flatten":    flatten,
	}
}

func RegisterFilter(name string, filter Filter) {
	filtersStore[name] = filter
}

func getFilter(name string) Filter {
	return filtersStore[name]
}
