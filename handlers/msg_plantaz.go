package handlers

import (
	"github.com/bwmarrin/discordgo"
	"padisoft/banana_farmer_bot/database"
	"strconv"
)

func (h *Handler) MsgPlantaz(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) {
	user, _ := db.GetUserData(m.Author.Username, m.Author.ID)

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x5f119e,
		Title:  m.Author.Username,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Vlastníš: " + strconv.Itoa(int(user["bananas"].(int32))) + " 🍌",
				Value:  "Miluju opice. 🐒 A taky banány!",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
