package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func embedPurple(title string, fieldName string, fieldValue string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x5f119e,
		Title:  title,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   fieldName,
				Value:  fieldValue,
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
		},
	}
}

func embedGold(title string, fields []*discordgo.MessageEmbedField) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x5f119e,
		Title:  title,
		Fields: fields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
		},
	}
}
