package logger

type Config struct {
	Debug bool `env:"DEBUG" envDefault:"false"`
}

type Option func(*Config)

func WithDebug(debug bool) Option {
	return func(c *Config) {
		c.Debug = debug
	}
}
