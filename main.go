package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/yokeTH/discord-bot/internal/bot"
	"github.com/yokeTH/discord-bot/internal/config"
	"github.com/yokeTH/discord-bot/pkg/logger"
)

func main() {
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	logger := logger.New(logger.WithDebug(cfg.Logger.Debug))

	myBot := bot.New(cfg.Bot, logger)
	myBot.Start(ctx, stop)
}
