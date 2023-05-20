package handlers

import (
	"github.com/bwmarrin/discordgo"
	"padisoft/banana_farmer_bot/database"
	"strconv"
)

func (h *Handler) MsgBTop(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) {
	topUsers, _ := db.GetTopUsers()

	var fields []*discordgo.MessageEmbedField
	//decodes the monkeys
	for i, monke := range topUsers {
		field := discordgo.MessageEmbedField{
			Name:   strconv.Itoa(i+1) + ". " + monke["userName"].(string),
			Value:  "Banánů: " + strconv.Itoa(int(monke["bananas"].(int32))),
			Inline: false,
		}
		fields = append(fields, &field)
	}
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0xfcba03, // Green
		Title:  "🐒** Nejlepší opičáci: **🐒",
		Fields: fields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
