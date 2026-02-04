package core

import (
	"fmt"
)

func CommandListChats(s *Session, args []string) error {
	i := 1
	for _, ch := range s.DB.Chats {
		chat, err := s.DB.GetByID(ch.ID)
		if err != nil {
			return fmt.Errorf("")
		}
		fmt.Printf("%d. %s\n", i, chat.Conversation[0].Content)
		i++
	}
	return nil
}
