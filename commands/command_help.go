package commands

import (
	"cli-chat/session"
	"fmt"
)

const (
	green = "\x1b[32m"
	reset = "\x1b[39m"
)

func CommandHelp(s *session.Session, args []string) error {
	for _, command := range GetCommands() {
		fmt.Printf("%s%s%s:%s %s\n", green, s.AppSettings.CommandPrefix, command.Name, reset, command.Description)
		fmt.Println()
	}
	return nil
}
