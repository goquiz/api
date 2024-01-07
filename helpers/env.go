package helpers

import (
	"flag"
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
	CookieDomain   string `env:"COOKIE_DOMAIN"`
}

type corsEnv struct {
	AllowOrigins string `env:"CORS_ALLOW_ORIGINS"`
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

type hCaptchaEnv struct {
	SiteKey   string `env:"HCAPTCHA_SITE_KEY"`
	SecretKey string `env:"HCAPTCHA_SECRET_KEY"`
}

type emailEnv struct {
	SMTPHost string `env:"EMAIL_SMTP_HOST"`
	SMTPPort int    `env:"EMAIL_SMTP_PORT"`
	Username string `env:"EMAIL_USERNAME"`
	Password string `env:"EMAIL_PASSWORD"`
}

type env struct {
	ServerPort string `env:"SERVER_PORT"`
	Session    sessionEnv
	Database   databaseEnv
	Redis      redisEnv
	Cors       corsEnv
	HCaptcha   hCaptchaEnv
	Email      emailEnv
}

var Env env

func (e *env) Load() {
	envPath := flag.String("env", ".env", "The environment file to load")
	flag.Parse()
	file, err := os.Open(*envPath)
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
