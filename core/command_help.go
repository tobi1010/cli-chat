package core

import (
	"fmt"
	"sort"
)

const (
	green = "\x1b[32m"
	reset = "\x1b[39m"
)

func CommandHelp(s *Session, args []string) error {
	metaList := s.Meta
	sort.Slice(metaList, func(i, j int) bool {
		return metaList[i].Name < metaList[j].Name
	})

	for _, m := range metaList {
		fmt.Printf("%s%s%s:%s %s\n", green, s.AppSettings.CommandPrefix, m.Name, reset, m.Description)
		fmt.Println()
	}
	return nil
}
