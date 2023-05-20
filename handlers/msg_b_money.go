package handlers

import (
	"github.com/bwmarrin/discordgo"
	"padisoft/banana_farmer_bot/database"
	"strconv"
)

func (h *Handler) MsgBMoney(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) {
	user, _ := db.GetUserData(m.Author.Username, m.Author.ID)
	money := 0
	if user["money"] != nil {
		money = int(user["money"].(int32))
	} else {
		db.AddMoney(m.Author.ID, 0)
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x5f119e,
		Title:  m.Author.Username,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Vlastníš: " + strconv.Itoa(money) + " Opicich dolaru",
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
