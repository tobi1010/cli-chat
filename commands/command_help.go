package commands

import (
	"fmt"
	"sort"
	"terminal-chat/session"
)

const (
	green = "\x1b[32m"
	reset = "\x1b[39m"
)

type Entry struct {
	Key string
	Val CliCommand
}

func CommandHelp(s *session.Session, args []string) error {
	commands := GetCommands()

	entries := make([]Entry, 0, len(commands))
	for k, v := range commands {
		entries = append(entries, Entry{Key: k, Val: v})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Val.Name < entries[j].Val.Name
	})
	for _, e := range entries {
		fmt.Printf("%s%s%s:%s %s\n", green, s.AppSettings.CommandPrefix, e.Val.Name, reset, e.Val.Description)
		fmt.Println()
	}
	return nil
}
