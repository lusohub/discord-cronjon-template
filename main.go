package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gtuk/discordwebhook"
)

func main() {
	discordURL := os.Getenv("DISCORD_URL")
	if discordURL == "" {
		log.Fatal("DISCORD_URL environment variable is required")
	}

	currentTime := time.Now().Format("15:04:05")
	content := fmt.Sprintf("Hora atual: %s", currentTime)
	username := "CronJob Bot"

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(discordURL, message)
	if err != nil {
		log.Fatal(err)
	}
}