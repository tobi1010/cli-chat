package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	cfg, err := newConfig("ChatGPT")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("using %s\n", cfg.model)
}
