package handlers

import (
	"github.com/bwmarrin/discordgo"
	"padisoft/banana_farmer_bot/database"
)

func (h *Handler) MsgBHovno(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) error {
	user, err := db.GetUserData(m.Author.Username, m.Author.ID)
	if err != nil {
		return err
	}

	if user["money"] == nil {
		err := db.AddMoney(m.Author.ID, 0)
		if err != nil {
			return err
		}
	}

	money := int(user["money"].(int32))
	if money > 100 {
		err := db.AddHovno(m.Author.ID)
		if err != nil {
			return err
		}

		err = db.AddMoney(m.Author.ID, -100)
		if err != nil {
			return err
		}

		embed := embedPurple(m.Author.Username, "Koupil/a/o jsi 1 opici hovno za 100 Opicich dolaru. Pouzij jej prikazem 'hovno @User'", "Miluju opice. 🐒 A taky banány!")

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)

		return err
	}

	_, err = s.ChannelMessageSendReply(m.ChannelID, "Potrebujes aspon 100 opicich dolaru pro koupi opiciho hovna", m.Reference())

	return err
}
