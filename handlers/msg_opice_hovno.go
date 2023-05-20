package handlers

import "github.com/bwmarrin/discordgo"

func (h *Handler) MsgOpiceHovno(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSendReply(m.ChannelID, "ZIJU TI VE ZDECH ZIJU TI VE ZDECH", m.Reference())
	return err
}
