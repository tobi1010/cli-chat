package session

import (
	"cli-chat/paths"
	"fmt"
)

func (s *Session) UpdateChat() error {
	if s.DB == nil {
		return fmt.Errorf("db is nil")
	}
	if s.Chat == nil {
		return fmt.Errorf("chat is nil")
	}
	chatsDir, err := paths.ChatsDir()
	if err != nil {
		return fmt.Errorf("resolving chats dir: %w", err)
	}
	if err := s.Chat.Write(chatsDir); err != nil {
		return fmt.Errorf("writing chat atomically: %w", err)
	}
	indexPath, err := paths.IndexPath()
	if err != nil {
		return fmt.Errorf("resolving index path: %w", err)
	}
	s.DB.Touch(s.Chat.ID, s.Chat.UpdatedAt)
	if err := s.DB.Save(indexPath); err != nil {
		return fmt.Errorf("saving db: %w", err)
	}
	return nil
}
