package command

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var ActivateCommand = &discordgo.ApplicationCommand{
	Name:        "activate",
	Description: "Activate another slash command globally",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:         discordgo.ApplicationCommandOptionString,
			Name:         "command",
			Description:  "Name of the command to activate (e.g., ping)",
			Required:     true,
			Autocomplete: true,
		},
	},
}

func AutocompleteHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommandAutocomplete {
		return
	}

	data := i.ApplicationCommandData()

	if data.Name == "activate" || data.Name == "deactivate" {
		input := data.Options[0].StringValue()
		var choices []*discordgo.ApplicationCommandOptionChoice

		knownCommands := []string{"ping"}

		for _, cmd := range knownCommands {
			if input == "" || startsWith(cmd, input) {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  cmd,
					Value: cmd,
				})
			}
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionApplicationCommandAutocompleteResult,
			Data: &discordgo.InteractionResponseData{
				Choices: choices,
			},
		})
	}
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HandleActivate(s *discordgo.Session, i *discordgo.InteractionCreate, appID string) {
	cmdName := i.ApplicationCommandData().Options[0].StringValue()
	var cmd *discordgo.ApplicationCommand

	switch cmdName {
	case "ping":
		cmd = PingCommand
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Unknown command: %s", cmdName),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	_, err := s.ApplicationCommandCreate(appID, "", cmd)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Failed to activate %s: %v", cmdName, err),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Activated command: %s", cmdName),
		},
	})
}
