package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	u "github.com/holy-tech/discord-roulette/src"
	h "github.com/holy-tech/discord-roulette/src/handlers"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	u.CheckErr("Invalid paramters: %v", err)

	discord.AddHandler(h.Ready)
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = discord.Open()
	u.CheckErr("Could not open session: %v", err)
	defer discord.Close()

	h.AppendHandler(discord, &h.RouletteHandle)
	h.AppendHandler(discord, &h.ShootHandle)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
