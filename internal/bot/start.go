package bot

import (
	"context"
	"fmt"

	"github.com/yokeTH/discord-bot/internal/command"
)

func (b *bot) Start(ctx context.Context, stop context.CancelFunc) {
	go func() {
		if err := b.session.Open(); err != nil {
			b.logger.Error(fmt.Sprintf("Cannot open Discord session: %v", err))
			stop()
		}
	}()

	b.session.AddHandler(command.CommandRouter)
	b.session.AddHandler(command.AutocompleteHandler)

	for _, cmd := range command.Commands {
		_, err := b.session.ApplicationCommandCreate(b.config.AppID, "", cmd)
		if err != nil {
			fmt.Printf("cannot create command %s: %v\n", cmd.Name, err)
		}
	}

	b.logger.Info("Bot is now running.")

	<-ctx.Done()

	b.logger.Info("Shutdown signal received. Cleaning up...")

	if err := b.shutdown(); err != nil {
		b.logger.Info(fmt.Sprintf("error during shutdown: %v", err))
	} else {
		b.logger.Info("bot shutdown completed.")
	}
}

func (b *bot) shutdown() error {
	if err := b.session.Close(); err != nil {
		return fmt.Errorf("failed to close Discord session: %w", err)
	}
	return nil
}
