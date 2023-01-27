package components

import (
	"os"

	env "github.com/joho/godotenv"
)

func init() {
	if _, err := os.Stat("./.env"); err == nil {
		if err := env.Load("./.env"); err != nil {
			panic(err.Error())
		}
	}

}

func EnvString(k string) string {
	return os.Getenv(k)
}
