package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type bot struct {
	session *discordgo.Session
	config  Config
	logger  *zap.Logger
}

func New(config Config, logger *zap.Logger) *bot {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(fmt.Sprintf("Error creating Discord session, %v", err))
	}

	return &bot{
		session: session,
		config:  config,
		logger:  logger,
	}
}
