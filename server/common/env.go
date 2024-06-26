package common

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Cannot laod environment file.")
	}
}
