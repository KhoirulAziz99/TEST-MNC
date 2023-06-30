package utility


import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//untuk keamanan data sensitive

func GetEnv(key string) string {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		 panic(err)
	}

	return os.Getenv(key)
}