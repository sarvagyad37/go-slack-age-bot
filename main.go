package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println("Timestamp: ", event.Timestamp)
		fmt.Println("Command: ", event.Command)
		fmt.Println("Parameters: ", event.Parameters)
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-xxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx")
	os.Setenv("SLACK_APP_TOKEN", "xapp-xxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my YOB is <year>", &slacker.CommandDefinition{
		Description: "YOB calculator",
		// Example: "my YOB is 1990",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				panic(err)
			}
			currentYear := time.Now().Year()
			age := currentYear - yob
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
