package command

import (
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	ActivateCommand,
	DeactivateCommand,
}

func CommandRouter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case ActivateCommand.Name:
		HandleActivate(s, i, s.State.User.ID)
	case DeactivateCommand.Name:
		HandleDeactivate(s, i, s.State.User.ID)
	case PingCommand.Name:
		HandlePing(s, i)
	}
}
