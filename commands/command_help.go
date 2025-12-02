package commands

import (
	"cli-chat/config"
	"fmt"
)

func CommandHelp(cfg *config.Config, args []string) error {
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}
