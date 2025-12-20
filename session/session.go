package session

import (
	"fmt"
)

func (s *Session) UpdateChat() error {
	if s.DB == nil {
		return fmt.Errorf("db is nil")
	}
	if s.Chat == nil {
		return fmt.Errorf("chat is nil")
	}
	if err := s.Chat.Write(s.Paths.ChatsDir); err != nil {
		return fmt.Errorf("writing chat atomically: %w", err)
	}
	s.DB.Touch(s.Chat.ID, s.Chat.UpdatedAt)
	if err := s.DB.Save(s.Paths.IndexPath); err != nil {
		return fmt.Errorf("saving db: %w", err)
	}
	return nil
}
