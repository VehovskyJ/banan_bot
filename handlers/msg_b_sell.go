package handlers

import (
	"github.com/bwmarrin/discordgo"
	"math"
	"padisoft/banana_farmer_bot/database"
	"strconv"
)

func (h *Handler) MsgBSell(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) {
	user, _ := db.GetUserData(m.Author.Username, m.Author.ID)
	bananas := int(user["bananas"].(int32))
	money := math.Round(float64(bananas / 5))
	db.ResetBananas(m.Author.ID, bananas)
	db.AddMoney(m.Author.ID, int(money))
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x5f119e,
		Title:  m.Author.Username,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Prodal/a/o jsi : " + strconv.Itoa(bananas) + "🍌 za " + strconv.Itoa(int(money)) + " Opicich dolaru",
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
