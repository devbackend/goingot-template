package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func MustLoad(filenames ...string) {
	err := godotenv.Load(filenames...)
	if err != nil {
		panic("Error loading .env file")
	}
}

func MustNotEmpty(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf(`expected env "%s" not empty!`, key))
	}

	return val
}
