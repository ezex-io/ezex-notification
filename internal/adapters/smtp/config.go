package smtp

type Config struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Pass      string `yaml:"pass"`
	FromEmail string `yaml:"from_email"`
}

func DefaultConfig() *Config {
	return &Config{
		Host:      "smtp.example.com",
		Port:      587,
		User:      "smtp_user",
		Pass:      "smtp_password",
		FromEmail: "no-reply@example.com",
	}
}

func (*Config) BasicCheck() error {
	return nil
}
