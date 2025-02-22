package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	
)

var (
	BotToken string
	GuildID string
	ClientID string
)

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(newMessage)

	discord.Open()
	defer discord.Close()

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	fmt.Println("Bot is running")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	

	switch {
	case strings.HasPrefix(message.Content, "/whois"):
		args := strings.Split(message.Content, " ")
		if len(args) == 1 {
			discord.ChannelMessageSend(message.ChannelID, "Please refer to a user")
			return
		}
		fmt.Println(args[1])
		discord.ChannelMessageSend(message.ChannelID, args[1])
	}
}