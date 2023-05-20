package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"padisoft/banana_farmer_bot/database"
	"time"
)

var scheduler gocron.Scheduler = gocron.Scheduler{}

func (h *Handler) MsgHovno(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) {
	scheduler = *gocron.NewScheduler(time.UTC)

	if len(m.Mentions) == 1 {
		res, _ := db.SubHovno(m.Author.Username, m.Author.ID)
		if !res {
			s.ChannelMessageSendReply(m.ChannelID, "Nemas dost hoven :// kup nejake pres b hovno", m.Reference())
		}
		scheduler.Clear()

		if m.Mentions[0].ID == "m.Mentions[0].ID" {
			s.ChannelMessageSend(m.ChannelID, "⚠️⚠️⚠️⚠️")
			s.ChannelMessageSend(m.ChannelID, "Za trest budes ohovnen")

			scheduler.Every(10).Seconds().Do(func() {
				s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> byl/a/o jsi proklet opicim prokletim. Hod hovno po nekom dalsim aby jsi se ho zbavil")
			},
			)
			scheduler.StartAsync()
		}

		s.ChannelMessageSendReply(m.ChannelID, "Hodil/a/o jsi opici hovno po <@"+m.Mentions[0].ID+">", m.Reference())

		scheduler.Every(2).Minutes().Do(func() {
			s.ChannelMessageSend(m.ChannelID, "<@"+m.Mentions[0].ID+"> byl/a/o jsi proklet opicim prokletim. Hod hovno po nekom dalsim aby jsi se ho zbavil")
		},
		)
		scheduler.StartAsync()
	}
}
