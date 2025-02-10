package main

import (
	"glchs_bot/bot"
	"log"
	"os"
	"fmt"
)

func main()  {
	data, err := os.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	botToken := string(data)
	bot.BotToken = botToken
	bot.Run()
}