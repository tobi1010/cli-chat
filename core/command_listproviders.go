package core

import (
	"fmt"
	"terminal-chat/providers"
)

func CommandListProviders(s *Session, args []string) error {
	i := 1
	for k := range providers.Registry {
		fmt.Printf("%d. %s\n", i, k)
		i++
	}
	return nil
}
