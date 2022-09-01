package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	handlers "github.com/holy-tech/discord-roulette/src"
)

func main() {
	discord, _ := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	discord.AddHandler(handlers.Ready)
	discord.AddHandler(handlers.Salute)

	err := discord.Open()
	defer discord.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	return
}
