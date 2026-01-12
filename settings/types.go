package settings

type Settings struct {
	DefaultProvider string `json:"default_provider"`
	DefaultModel    string `json:"default_model"`
	CommandPrefix   string `json:"command_prefix"`
	Timeout         int    `json:"timeout"`
	Columns         int    `json:"columns"`
	ApiKey          string `json:"api_key"`
	TTL             int    `json:"ttl"`
}

func NewDefaultSettings() Settings {
	return Settings{
		DefaultProvider: "openai",
		DefaultModel:    "gpt-5-nano",
		CommandPrefix:   "/",
		Timeout:         60,
		Columns:         80,
		ApiKey:          "OPENAI_API_KEY",
		TTL:             60 * 60 * 24 * 7, // 1 week
	}
}
