package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	handlers "github.com/holy-tech/discord-roulette/src"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid paramters: %v", err)
	}

	discord.AddHandler(handlers.Ready)

	if err := discord.Open(); err != nil {
		log.Fatalf("Could not open session: %v", err)
	}
	defer discord.Close()

	handlers.AppendHandler(discord, &handlers.RouletteHandle)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
