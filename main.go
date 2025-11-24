package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	s, err := readSettings(".")
	cfg, err := newConfig(s)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("using %s\n", cfg.Settings.Model)
	startRepl(&cfg)
}
