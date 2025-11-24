package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, arg string) error {
	fmt.Println("Goodbye")
	os.Exit(0)
	return fmt.Errorf("Error closing chat!")
}
