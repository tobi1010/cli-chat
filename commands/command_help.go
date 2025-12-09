package commands

import (
	"cli-chat/session"
	"fmt"
)

func CommandHelp(s *session.Session, args []string) error {
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}
