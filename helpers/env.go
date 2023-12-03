package helpers

import (
	"fmt"
	"github.com/golobby/dotenv"
	"os"
)

type env struct {
	ServerPort string `env:"SERVER_PORT"`
	JwtSecret  string `env:"JWT_SECRET"`
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
