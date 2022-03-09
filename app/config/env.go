package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error in loading the .env file")
    }
  
    return os.Getenv("MONGOURI")
}