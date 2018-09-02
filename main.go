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
	message := fmt.Sprintf(
		"My message with a new line %v break",
		builder.AddLineBreak())

	// Send Message
	builder.SendMessage(message)

}
