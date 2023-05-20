package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"padisoft/banana_farmer_bot/database"
)

func (h *Handler) MsgPlantaz(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) error {
	user, err := db.GetUserData(m.Author.Username, m.Author.ID)
	if err != nil {
		return err
	}

	embed := embedPurple(m.Author.Username, fmt.Sprintf("Vlastníš: %d 🍌", int(user["bananas"].(int32))), "Miluju opice. 🐒 A taky banány!")

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)

	return err
}
