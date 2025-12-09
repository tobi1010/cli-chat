package settings

type Settings struct {
	CommandPrefix string `json:"command_prefix"`
	Timeout       int    `json:"timeout"`
	Columns       int    `json:"columns"`
}

func NewDefaultSettings() Settings {
	return Settings{
		CommandPrefix: "/",
		Timeout:       60,
		Columns:       80,
	}
}
