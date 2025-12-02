package commands

import (
	"cli-chat/config"
	"fmt"
	"os"
)

func CommandExit(cfg *config.Config, args []string) error {
	fmt.Println("Goodbye")
	os.Exit(0)
	return fmt.Errorf("Error closing chat!")
}
