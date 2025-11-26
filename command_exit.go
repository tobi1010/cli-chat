package main

import (
	"cli-chat/config"
	"fmt"
	"os"
)

func commandExit(cfg *config.Config, arg string) error {
	fmt.Println("Goodbye")
	os.Exit(0)
	return fmt.Errorf("Error closing chat!")
}
