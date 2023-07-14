package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Krognol/go-wolfram"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

var wolframClient *wolfram.Client

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	slack_bot_token := os.Getenv("SLACK_BOT_TOKEN")
	slack_app_token := os.Getenv("SLACK_APP_TOKEN")
	wolfram_app_id := os.Getenv("WOLFRAM_APP_ID")

	bot := slacker.NewClient(slack_bot_token, slack_app_token)
	client := witai.NewClient(wolfram_app_id)
	wolframClient = &wolfram.Client{AppID: wolfram_app_id}

	go printCommandEvents(bot.CommandEvents())

	bot.Command("query for bot - <message>", &slacker.CommandDefinition{
		Description: "send any question to wolfrm",
		Examples:    []string{"who is the president of usa"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")
			fmt.Println(query)
			msg, err := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			if err != nil {
				log.Println("Error parsing message:", err)
				// Handle the error appropriately
				response.Reply("Oops! There was an error processing your request.")
				return
			}
			data, _ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			answer := value.String()
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				fmt.Println("there is an error")
			}
			fmt.Println(res)
			response.Reply("received")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
