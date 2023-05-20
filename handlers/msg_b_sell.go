package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math"
	"padisoft/banana_farmer_bot/database"
)

func (h *Handler) MsgBSell(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) error {
	user, err := db.GetUserData(m.Author.Username, m.Author.ID)
	if err != nil {
		return err
	}

	bananas := int(user["bananas"].(int32))
	money := int(math.Round(float64(bananas / 5)))

	err = db.ResetBananas(m.Author.ID, bananas)
	if err != nil {
		return err
	}

	err = db.AddMoney(m.Author.ID, money)
	if err != nil {
		return err
	}

	embed := embedPurple(m.Author.Username, fmt.Sprintf("Prodal/a/o jsi : %d🍌 za %d Opicich dolaru", bananas, money), "Miluju opice. 🐒 A taky banány!")

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)

	return err
}
