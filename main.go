package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"padisoft/banana_farmer_bot/database"
	"padisoft/banana_farmer_bot/handlers"
	"strings"
	"syscall"
	//"time"

	"github.com/bwmarrin/discordgo"
)

// var token = os.Getenv("DISCORD_TOKEN")
var token = "MTA5MjUzMzMzMjU1MTY4NDIyOA.GS3Sbk.maOBUBFWs0Wu2sqRWGLZFaRNegy9ks5yALGsjE"

func main() {
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

	handler := handlers.NewHandler(s, m, db)

	switch strings.ToLower(m.Content) {
	case "b":
		handler.MsgB(s, m, db)
	case "plantaz":
		handler.MsgPlantaz(s, m, db)
	case "b money":
		handler.MsgBMoney(s, m, db)
	case "b sell":
		handler.MsgBSell(s, m, db)
	case "b hovno":
		handler.MsgBHovno(s, m, db)
	case "hovno":
		handler.MsgHovno(s, m, db)
	case "b top":
		handler.MsgBTop(s, m, db)
	case "get nerded":
		handler.MsgGetNerded(s, m)
	case "get jinxed":
		handler.MsgGetJinxed(s, m)
	case "opice hovno":
		handler.MsgOpiceHovno(s, m)
	}
}
