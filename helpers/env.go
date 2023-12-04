package helpers

import (
	"fmt"
	"github.com/golobby/dotenv"
	"os"
)

type sessionEnv struct {
	ExpiresIn      int    `env:"SESSION_EXPIRES_IN"`
	KeyLookup      string `env:"SESSION_KEY_LOOKUP"`
	CookieSecure   bool   `env:"COOKIE_SECURE"`
	CookieHttpOnly bool   `env:"COOKIE_HTTP_ONLY"`
	CookieSameSite string `env:"COOKIE_SAME_SITE"`
}

type databaseEnv struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Database string `env:"DB_NAME"`
}

type redisEnv struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	Username string `env:"REDIS_USERNAME"`
	Password string `env:"REDIS_PASSWORD"`
}

type env struct {
	ServerPort string `env:"SERVER_PORT"`
	JwtSecret  string `env:"JWT_SECRET"`
	Session    sessionEnv
	Database   databaseEnv
	Redis      redisEnv
}

var Env env

func (e *env) Load() {
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println("Failed to load environment variables:", err)
		os.Exit(1)
	}
	err = dotenv.NewDecoder(file).Decode(&*e)
	if err != nil {
		fmt.Println("Failed to parse environment variables:", err)
		os.Exit(1)
	}
}
