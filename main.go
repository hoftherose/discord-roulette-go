package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateGameStatus(0, "!roulette")
}

func Salute(session *discordgo.Session, event *discordgo.MessageCreate) {
	fmt.Println(event.Message)
}

func main() {
	discord, _ := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	discord.AddHandler(Ready)
	fmt.Println("Ready")
	discord.AddHandler(Salute)
	fmt.Println("Salute")

	err := discord.Open()
	defer discord.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("Dying")

	return
}
