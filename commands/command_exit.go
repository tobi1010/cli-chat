package commands

import (
	"errors"
	"fmt"
	"terminal-chat/session"
)

var ErrExitRequested = errors.New("exit requested")

func CommandExit(s *session.Session, args []string) error {
	fmt.Printf("\nGoodbye\n")
	return ErrExitRequested
}
