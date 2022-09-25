package embed

import (
	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
)

func TempEmbed(s *discordgo.Session, channel string) {
	e := embed.NewEmbed()
	e.SetTitle("FakeTitle")
	s.ChannelMessageSendEmbed(channel, e.MessageEmbed)
	s.ChannelMessageSendEmbed(channel, embed.NewGenericEmbed("Example", "This is an example embed!"))
	s.ChannelMessageSendEmbed(channel, embed.NewErrorEmbed("Example Error", "This is an example error embed!"))
}
