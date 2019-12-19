package config

// DB holds the configuration values for the database.
type DB struct {
	Host     string `env:"DB_HOST" default:"127.0.0.1"`
	Port     int    `env:"DB_PORT" default:"3306"`
	User     string `env:"DB_USER" default:"root"`
	Password string `env:"DB_PASSWORD" default:"Kolian18**"`
	Name     string `env:"DB_NAME" default:"municipality"`
	Limit    int    `env:"DB_PAGE_LIMIT" default:"50"`
}