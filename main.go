package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"strconv"
	"strings"
	"syscall"

	//"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	//flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

var token string = "OTg0NzczMDE0NTAxNjE3Njg0.GNEQyQ.5ekJcE4BKChZdGsJ6GyOC7wJsnyBa1B_zRrvcM"
var dbClient mongo.Client = mongo.Client{}

func main() {

	//Init banana database
	dbClient = *initDatabase()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
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
	s.UpdateGameStatus(0, "epicka bananova plantáž")

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.ToLower(m.Content) == "b" {
		_ = GetUserData(dbClient, m.Author.ID)
		banans := rand.Intn(16)

		addBanans(dbClient, m.Author.ID, banans)

		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0x5f119e,
			Title:  "🐒Získal/a jsi banány!🐒",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Dostal/a jsi: " + strconv.Itoa(int(banans)) + " 🍌",
					Value:  "Miluju banány. A taky opice. ",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Credits: @Matyslav_  ||  Přispěj na vývoj opičáka na patreon.com/Padisoft 🐒",
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)

	}

}
func initDatabase() *mongo.Client {
	//MongoDB databse connection
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://monkiopicak:JB5NR5RJImwhLxtN@monkidatabse.cxodm.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
func GetUserData(client mongo.Client, userId string) bson.M {
	collection := client.Database("farmsDb").Collection("userFarm")
	var opicak bson.M
	if err := collection.FindOne(context.TODO(), bson.M{"userId": userId}).Decode(&opicak); err != nil {
		log.Print(err)
	} else {
		collection := client.Database("serversDb").Collection("servers")
		_, err := collection.InsertOne(context.TODO(), bson.D{{"userId", userId}, {"bananas", 0}, {"xp", 0}})
		if err != nil {

		}
	}
	return opicak
}
func addBanans(client mongo.Client, userId string, banans int) {
	collection := client.Database("farmsDb").Collection("userFarm")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"userId": userId},
		bson.D{
			{"$inc", bson.D{{"bananas", 1}}},
		},
	)
	if err != nil {
		log.Print(err)
	}
}
