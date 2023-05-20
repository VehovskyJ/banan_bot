package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"padisoft/banana_farmer_bot/database"
)

func (h *Handler) MsgB(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) error {
	bananas := rand.Intn(16)

	err := db.AddBananas(m.Author.ID, bananas)
	if err != nil {
		return err
	}

	embed := embedPurple(m.Author.Username, fmt.Sprintf("Dostal/a jsi: %d 🍌", bananas), "🐒Získal/a jsi banány!🐒")
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)

	return err
}
