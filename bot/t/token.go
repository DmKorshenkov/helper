package t

import (
	"os"

	"github.com/joho/godotenv"
)

func Token() string {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	return os.Getenv("dimocracy")
}
