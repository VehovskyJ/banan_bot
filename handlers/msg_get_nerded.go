package handlers

import "github.com/bwmarrin/discordgo"

func (h *Handler) MsgGetNerded(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.Mentions) == 1 {
		nerdedMessagesCount := 0
		nerdedMessagesTBD := 15
		messages, _ := s.ChannelMessages(m.ChannelID, 100, "", "", "")
		for _, message := range messages {
			if message.Author.ID == m.Mentions[0].ID {
				s.MessageReactionAdd(message.ChannelID, message.ID, "🤓")
				nerdedMessagesCount++
				nerdedMessagesTBD--
				if nerdedMessagesTBD == 0 {
					break
				}
			}
		}

	}
}
