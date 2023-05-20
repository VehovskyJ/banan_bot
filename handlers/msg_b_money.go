package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"padisoft/banana_farmer_bot/database"
)

func (h *Handler) MsgBMoney(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) error {
	user, err := db.GetUserData(m.Author.Username, m.Author.ID)
	if err != nil {
		return err
	}

	if user["money"] == nil {
		err = db.AddMoney(m.Author.ID, 0)
		if err != nil {
			return err
		}
	}

	money := int(user["money"].(int32))

	embed := embedPurple(m.Author.Username, fmt.Sprintf("Vlastníš: %d Opicich dolaru", money), "Miluju opice. 🐒 A taky banány!")

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)

	return err
}
