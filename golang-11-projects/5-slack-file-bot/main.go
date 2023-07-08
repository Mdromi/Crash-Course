package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	bootToken := os.Getenv("SLACK_BOOT_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")

	api := slack.New(bootToken)
	channelArr := []string{channelID}
	fileArr := []string{"sql.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name %s, URL:%s\n", file.Name, file.URL)
	}
}
