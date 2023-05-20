package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"padisoft/banana_farmer_bot/database"
	"strconv"
	"strings"
	"syscall"
	"time"

	//"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

var token = os.Getenv("DISCORD_TOKEN")
var scheduler gocron.Scheduler = gocron.Scheduler{}

func main() {
	scheduler = *gocron.NewScheduler(time.UTC)

	//Init banana database
	dbClient, err := database.Connect("mongodb+srv://monkiopicak:JB5NR5RJImwhLxtN@monkidatabse.cxodm.mongodb.net/?retryWrites=true&w=majority")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %s", err)
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		messageCreate(s, m, dbClient)
	})

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %s", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Delam opici zvuky.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	//s.UpdateGameStatus(0, "Krsipina smrdi jak tvoje mamka lool opice")
	err := s.UpdateGameStatus(0, "epicka bananova plantáž")
	if err != nil {
		log.Printf("Failed to set status: %s", err)
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) {
	if m.Author.Bot {
		return
	}

	switch strings.ToLower(m.Content) {
	case "b":
		_, _ = db.GetUserData(m.Author.Username, m.Author.ID)
		banans := rand.Intn(16)

		db.AddBananas(m.Author.ID, banans)

		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0x5f119e,
			Title:  m.Author.Username,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Dostal/a jsi: " + strconv.Itoa(int(banans)) + " 🍌",
					Value:  "🐒Získal/a jsi banány!🐒",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	case "plantaz":
		user, _ := db.GetUserData(m.Author.Username, m.Author.ID)

		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0x5f119e,
			Title:  m.Author.Username,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Vlastníš: " + strconv.Itoa(int(user["bananas"].(int32))) + " 🍌",
					Value:  "Miluju opice. 🐒 A taky banány!",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	case "b money":
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
	case "b sell":
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
	case "b hovno":
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
	case "hovno":
		if len(m.Mentions) == 1 {

			res, _ := db.SubHovno(m.Author.Username, m.Author.ID)
			if !res {
				s.ChannelMessageSendReply(m.ChannelID, "Nemas dost hoven :// kup nejake pres b hovno", m.Reference())
				return
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

				return
			}

			s.ChannelMessageSendReply(m.ChannelID, "Hodil/a/o jsi opici hovno po <@"+m.Mentions[0].ID+">", m.Reference())

			scheduler.Every(2).Minutes().Do(func() {
				s.ChannelMessageSend(m.ChannelID, "<@"+m.Mentions[0].ID+"> byl/a/o jsi proklet opicim prokletim. Hod hovno po nekom dalsim aby jsi se ho zbavil")
			},
			)
			scheduler.StartAsync()

		}
	case "b top":
		topUsers, _ := db.GetTopUsers()

		var fields []*discordgo.MessageEmbedField
		//decodes the monkeys
		for i, monke := range topUsers {
			field := discordgo.MessageEmbedField{
				Name:   strconv.Itoa(i+1) + ". " + monke["userName"].(string),
				Value:  "Banánů: " + strconv.Itoa(int(monke["bananas"].(int32))),
				Inline: false,
			}
			fields = append(fields, &field)
		}
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0xfcba03, // Green
			Title:  "🐒** Nejlepší opičáci: **🐒",
			Fields: fields,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	case "get nerded":
		if len(m.Mentions) == 1 {
			nerdedMessagesCount := 0
			nerdedMessagesTBD := 15
			messages, _ := s.ChannelMessages(m.ChannelID, 100, "", "", "")
			for _, message := range messages {
				if message.Author.ID == m.Mentions[0].ID {
					s.MessageReactionAdd(message.ChannelID, message.ID, "🤓")
					nerdedMessagesCount++
					nerdedMessagesTBD--
					if nerdedMessagesTBD == 0 {
						break
					}
				}
			}

		}
	case "get jinxed":
		if len(m.Mentions) == 1 {
			nerdedMessagesCount := 0
			nerdedMessagesTBD := 15
			messages, _ := s.ChannelMessages(m.ChannelID, 100, "", "", "")
			for _, message := range messages {
				if message.Author.ID == m.Mentions[0].ID {
					s.MessageReactionAdd(message.ChannelID, message.ID, "jinx1:1074460008307245067")
					nerdedMessagesCount++
					nerdedMessagesTBD--
					if nerdedMessagesTBD == 0 {
						break
					}
				}
			}

		}
	case "opice hovno":
		s.ChannelMessageSendReply(m.ChannelID, "ZIJU TI VE ZDECH ZIJU TI VE ZDECH", m.Reference())
	}
}
