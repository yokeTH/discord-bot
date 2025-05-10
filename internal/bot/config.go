package bot

type Config struct {
	Token   string `env:"BOT_TOKEN,required"`
	OwnerID string `env:"OWNER_ID,required"`
}
