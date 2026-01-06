package commands

import (
	"cli-chat/providers"
	"cli-chat/session"
	"fmt"
)

func CommandListProviders(s *session.Session, args []string) error {
	i := 1
	for k := range providers.Registry {
		fmt.Printf("%d. %s\n", i, k)
		i++
	}
	return nil
}
