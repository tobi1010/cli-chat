package core

import (
	"fmt"
)

func CommandPrintSettings(s *Session, args []string) error {
	set := s.AppSettings
	if err := set.PrintSettings(); err != nil {
		return fmt.Errorf("printig settings: %w", err)
	}
	return nil
}
