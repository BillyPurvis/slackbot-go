package main

import (
	"fmt"

	"github.com/BillyPurvis/go-slack-bot/builder"
)

// Slack Incoming Webhook URL

// Message A Go Message
type Message struct {
	Text string
}

func main() {

	fmt.Println("Waking up Isobel")

	// Create Message In App
	message := builder.CreateMessage("Hello World")

	builder.AddButtons(
		message,
		"Would you like to play a game, Sir?",
		builder.Actions{
			builder.Action{
				"Game",
				"Test",
				"button",
				"test",
			},
			builder.Action{
				"Game",
				"Test",
				"button",
				"test",
			},
		},
	)

	// //TODO Add the buttons to a slice [] using append
	// builder.AddButton(message, "Would you like to play a game?", builder.Action{
	// 	"Game",
	// 	"Test",
	// 	"button",
	// 	"test",
	// },
	// )
	// builder.AddButton(message, "Would you like to play a game?", builder.Action{
	// 	"Game",
	// 	"Test",
	// 	"button",
	// 	"test",
	// },

	// Send the message
	builder.SendMessage(message)

}
