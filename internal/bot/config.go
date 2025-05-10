package bot

type Config struct {
	AppID   string `env:"APP_ID,required"`
	Token   string `env:"BOT_TOKEN,required"`
	OwnerID string `env:"OWNER_ID,required"`
}
