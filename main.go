package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token string
var prefix = "!"

func init() {
	flag.StringVar(&Token, "t", "", "The token for your bot")
	flag.Parse()
}

func main() {
	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("[ERROR]: creating Discord session,", err)
		return
	}
	bot.AddHandler(messageCreate)
	err = bot.Open()
	if err != nil {
		fmt.Println("[ERROR]: opening connection,", err)
		return
	}
	fmt.Println("[BOT]: Running     Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
}

func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	if !strings.HasPrefix(msg.Content, prefix) {
		return
	}
  command:= strings.Trim(msg.Content, prefix)
	if command == "ping" {
		session.ChannelMessageSend(msg.ChannelID, "pong")
	}

	if command == "test-embed" {
		embed := &discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Test Embed",
			Description: "Ini descripsinya",
		}
		session.ChannelMessageSendEmbed(msg.ChannelID, embed)
	}
}
