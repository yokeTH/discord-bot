package command

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var DeactivateCommand = &discordgo.ApplicationCommand{
	Name:        "deactivate",
	Description: "Remove a global command by name",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:         discordgo.ApplicationCommandOptionString,
			Name:         "command",
			Description:  "Command name to deactivate",
			Required:     true,
			Autocomplete: true,
		},
	},
}

func HandleDeactivate(s *discordgo.Session, i *discordgo.InteractionCreate, appID string) {
	cmdName := i.ApplicationCommandData().Options[0].StringValue()

	cmds, err := s.ApplicationCommands(appID, "")
	if err != nil {
		respondEphemeral(s, i, fmt.Sprintf("Failed to fetch global commands: %v", err))
		return
	}

	for _, cmd := range cmds {
		if cmd.Name == cmdName {
			err := s.ApplicationCommandDelete(appID, "", cmd.ID)
			if err != nil {
				respondEphemeral(s, i, fmt.Sprintf("Failed to delete command %s: %v", cmdName, err))
			} else {
				respondText(s, i, fmt.Sprintf("Command `%s` deactivated globally.", cmdName))
			}
			return
		}
	}

	respondEphemeral(s, i, fmt.Sprintf("Command `%s` not found among global commands.", cmdName))
}

func respondText(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
}

func respondEphemeral(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
