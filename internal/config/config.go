package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/yokeTH/discord-bot/internal/bot"
	"github.com/yokeTH/discord-bot/pkg/logger"
)

type config struct {
	Bot    bot.Config    `envPrefix:"DISCORD_"`
	Logger logger.Config `envPrefix:"LOGGER_"`
}

func Load() *config {
	config := &config{}

	if err := godotenv.Load(); err != nil {
		fmt.Println("Unable to load .env file:", err)
	}

	if err := env.Parse(config); err != nil {
		panic(fmt.Sprintf("Unable to parse env vars: %s", err))
	}

	return config
}
