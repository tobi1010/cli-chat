package commands

import (
	"cli-chat/session"
	"errors"
	"fmt"
)

var ErrExitRequested = errors.New("exit requested")

func CommandExit(s *session.Session, args []string) error {
	fmt.Println("Goodbye")
	return ErrExitRequested
}
