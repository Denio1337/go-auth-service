package config

type EnvKey string

// Env params keys
const (
	EnvAppAddress EnvKey = "APP_ADDRESS"
	EnvDBPort     EnvKey = "DB_PORT"
	EnvDBUser     EnvKey = "DB_USER"
	EnvDBPassword EnvKey = "DB_PASSWORD"
	EnvDBName     EnvKey = "DB_NAME"
	EnvSecret     EnvKey = "SECRET"
	EnvWebhookURL EnvKey = "WEBHOOK_URL"
)

// Check env key validity
func (e EnvKey) IsValid() bool {
	switch e {
	case EnvAppAddress, EnvDBPort, EnvDBUser,
		EnvDBPassword, EnvDBName, EnvSecret, EnvWebhookURL:
		return true
	default:
		return false
	}
}
