package settings

type Settings struct {
	CommandPrefix string `json:"command_prefix"`
	Timeout       int    `json:"timeout"`
	Columns       int    `json:"columns"`
	ApiKey        string `json:"api_key"`
	TTL           int
}

func NewDefaultSettings() Settings {
	return Settings{
		CommandPrefix: "/",
		Timeout:       60,
		Columns:       80,
		ApiKey:        "OPENAI_API_KEY",
		TTL:           60 * 60 * 24 * 7, // 1 week
	}
}
