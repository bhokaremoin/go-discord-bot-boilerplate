package main

import (
	"fmt"

	"github.com/bhokaremoin/go-discord-bot-boilerplate/bot"
	"github.com/bhokaremoin/go-discord-bot-boilerplate/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error While Reading Config File, Error := %v", err)
	}

	bot.Start()
	<-make(chan struct{})
}
