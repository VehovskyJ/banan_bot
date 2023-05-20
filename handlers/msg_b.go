package handlers

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"padisoft/banana_farmer_bot/database"
	"strconv"
)

func (h *Handler) MsgB(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) {
	_, _ = db.GetUserData(m.Author.Username, m.Author.ID)
	banans := rand.Intn(16)

	db.AddBananas(m.Author.ID, banans)

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x5f119e,
		Title:  m.Author.Username,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Dostal/a jsi: " + strconv.Itoa(int(banans)) + " 🍌",
				Value:  "🐒Získal/a jsi banány!🐒",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
