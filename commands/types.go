package commands

import (
	"cli-chat/session"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*session.Session, []string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "exit application",
			Callback:    CommandExit,
		},
		"print-settings": {
			Name:        "print-settings",
			Description: "print current settings",
			Callback:    CommandPrintSettings,
		},
		"set-prefix": {
			Name:        "set-prefix",
			Description: "set command prefix to <char>",
			Callback:    CommandSetPrefix,
		},
		"set-columns": {
			Name:        "set-columns",
			Description: "set line wrap in columns",
			Callback:    CommandSetColumns,
		},
		"list-chats": {
			Name:        "list-chats",
			Description: "",
			Callback:    CommandListChats,
		},
		"list-models": {
			Name:        "list-models",
			Description: "lists all available models of current provider",
			Callback:    CommandListModels,
		},
		"list-providers": {
			Name:        "list-providers",
			Description: "lists all available providers",
			Callback:    CommandListProviders,
		},
		"switch-provider": {
			Name:        "switch-provider",
			Description: "switch ai provider",
			Callback:    CommandSwitchProvider,
		},
		"switch-model": {
			Name:        "switch-model",
			Description: "switch llm",
			Callback:    CommandSwitchModel,
		},
		"usage": {
			Name:        "usage",
			Description: "",
			Callback:    CommandUsage,
		},
		"help": {
			Name:        "help",
			Description: "print help",
			Callback:    CommandHelp,
		},
		"new": {
			Name:        "new",
			Description: "start new chat",
			Callback:    CommandNew,
		},
		"set-apikey": {
			Name:        "set-apikey",
			Description: "set api key for provider",
			Callback:    CommandSetApiKey,
		},
	}
}
