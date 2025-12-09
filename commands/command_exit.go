package commands

import (
	"cli-chat/session"
	"fmt"
	"os"
)

func CommandExit(s *session.Session, args []string) error {
	fmt.Println("Goodbye")
	os.Exit(0)
	return fmt.Errorf("Error closing chat!")
}
