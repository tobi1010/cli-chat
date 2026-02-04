package core

import "fmt"

type CliCommand struct {
	Name        string
	Aliases     []string
	Description string
	Callback    func(*Session, []string) error
}

type CommandMeta struct {
	Name        string
	Aliases     []string
	Description string
}

type CommandRegistry map[string]CliCommand

var commands = []CliCommand{
	{
		Name:        "exit",
		Aliases:     []string{"q"},
		Description: "exit application",
		Callback:    CommandExit,
	},
	{
		Name:        "print-settings",
		Aliases:     []string{"ps"},
		Description: "print current settings",
		Callback:    CommandPrintSettings,
	},
	{
		Name:        "set-prefix",
		Aliases:     []string{"spfx"},
		Description: "set command prefix to <char>",
		Callback:    CommandSetPrefix,
	},
	{
		Name:        "set-columns",
		Aliases:     []string{"sc"},
		Description: "set line wrap in columns",
		Callback:    CommandSetColumns,
	},
	{
		Name:        "list-chats",
		Aliases:     []string{"lc"},
		Description: "list stored chats",
		Callback:    CommandListChats,
	},
	{
		Name:        "list-models",
		Aliases:     []string{"lm"},
		Description: "lists all available models of current provider",
		Callback:    CommandListModels,
	},
	{
		Name:        "list-providers",
		Aliases:     []string{"lp"},
		Description: "lists all available providers",
		Callback:    CommandListProviders,
	},
	{
		Name:        "switch-provider",
		Aliases:     []string{"sp"},
		Description: "switch ai provider",
		Callback:    CommandSwitchProvider,
	},
	{
		Name:        "switch-model",
		Aliases:     []string{"sm"},
		Description: "switch llm",
		Callback:    CommandSwitchModel,
	},
	{
		Name:        "usage",
		Aliases:     []string{"u"},
		Description: "not implemented",
		Callback:    CommandUsage,
	},
	{
		Name:        "help",
		Aliases:     []string{"h"},
		Description: "print help",
		Callback:    CommandHelp,
	},
	{
		Name:        "new",
		Aliases:     []string{"n"},
		Description: "start new chat",
		Callback:    CommandNew,
	},
	{
		Name:        "set-apikey",
		Aliases:     []string{"sk"},
		Description: "set api key for provider",
		Callback:    CommandSetApiKey,
	},
	{
		Name:        "set-ttl",
		Aliases:     []string{"sttl"},
		Description: "set time to live for cahce",
		Callback:    CommandSetTTL,
	},
	{
		Name:        "set-default-model",
		Aliases:     []string{"sdm"},
		Description: "set the default model",
		Callback:    CommandSetDefautlModel,
	},
	{
		Name:        "set-default-provider",
		Aliases:     []string{"sdp"},
		Description: "set the default provider",
		Callback:    CommandSetDefautlProvider,
	},
}

func NewRegistry() (CommandRegistry, error) {
	reg := make(map[string]CliCommand)
	for _, cmd := range commands {
		reg[cmd.Name] = cmd
		for _, alias := range cmd.Aliases {
			if _, ok := reg[alias]; !ok {
				reg[alias] = cmd
			} else {
				return CommandRegistry{}, fmt.Errorf("duplicate alias. please edit settings.json again")
			}
		}
	}
	return reg, nil
}

func NewCommandMeta() []CommandMeta {
	meta := make([]CommandMeta, len(commands))
	for i, cmd := range commands {
		meta[i] = CommandMeta{Name: cmd.Name, Aliases: cmd.Aliases, Description: cmd.Description}
	}
	return meta
}
