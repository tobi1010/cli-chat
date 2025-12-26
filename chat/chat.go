package chat

import "time"

func (c *Chat) AddMessage(msg Message) {
	c.Conversation = append(c.Conversation, msg)
	c.UpdatedAt = time.Now()
}
