package core

import (
	"errors"
	"fmt"
)

var ErrExitRequested = errors.New("exit requested")

func CommandExit(s *Session, args []string) error {
	fmt.Printf("\nGoodbye\n")
	return ErrExitRequested
}
