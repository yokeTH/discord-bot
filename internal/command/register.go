package command

import (
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	PingCommand,
}

func CommandRouter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case PingCommand.Name:
		HandlePing(s, i)
	}
}
