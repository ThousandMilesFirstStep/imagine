package domain

type ConfigVips struct {
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
	Vips    ConfigVips
}
