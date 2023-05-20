package handlers

import (
	"github.com/bwmarrin/discordgo"
	"padisoft/banana_farmer_bot/database"
)

func (h *Handler) MsgBHovno(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) {
	user, _ := db.GetUserData(m.Author.Username, m.Author.ID)
	if user["money"] == nil {
		db.AddMoney(m.Author.ID, 0)
	}
	money := int(user["money"].(int32))
	if money > 100 {
		db.AddHovno(m.Author.ID)
		db.AddMoney(m.Author.ID, -100)
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0x5f119e,
			Title:  m.Author.Username,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Koupil/a/o jsi 1 opici hovno za 100 Opicich dolaru. Pouzij jej prikazem 'hovno @User'",
					Value:  "Miluju opice. 🐒 A taky banány!",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	} else {
		s.ChannelMessageSendReply(m.ChannelID, "Potrebujes aspon 100 opicich dolaru pro koupi opiciho hovna", m.Reference())
	}
}
