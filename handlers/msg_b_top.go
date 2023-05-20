package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"padisoft/banana_farmer_bot/database"
)

func (h *Handler) MsgBTop(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) error {
	topUsers, err := db.GetTopUsers()
	if err != nil {
		return err
	}

	var fields []*discordgo.MessageEmbedField
	for i, monke := range topUsers {
		field := discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%d. %s", i+1, monke["userName"].(string)),
			Value:  fmt.Sprintf("Banánů: %d", int(monke["bananas"].(int32))),
			Inline: false,
		}
		fields = append(fields, &field)
	}

	embed := embedGold("🐒** Nejlepší opičáci: **🐒", fields)

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)

	return err
}
