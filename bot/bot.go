package bot

import (
	"fmt"

	"github.com/bhokaremoin/go-discord-bot-boilerplate/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string

// var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Printf("Error := %v", err)
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Printf("Error := %v", err)
	}
	BotID = u.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Printf("Error := %v", err)
	}
	fmt.Println("Bot is Running")
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong")
	}
}
