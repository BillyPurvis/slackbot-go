package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	// Autoload for loading .ENV file
	_ "github.com/joho/godotenv/autoload"
)

// MessageStruct A message struct
type MessageStruct struct {
	Text        string `json:"text"`
	Attachments `json:"attachments,omitempty"`
}

// Attachment Make Attachment
type Attachment struct {
	Text    string `json:"text"`
	Actions `json:"actions"`
}

// Attachments Slice of Attachment
type Attachments []Attachment

// Action Set Action button options
type Action struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Actions Slack Actions
type Actions []Action

// CreateMessage Create a Message
func CreateMessage(message string) *MessageStruct {

	// Create new Message Struct
	messageStruct := &MessageStruct{
		message,
		nil,
	}
	return messageStruct
}

// AddButtons Adds Button to slack message
func AddButtons(message *MessageStruct, attachmentMessage string, actions Actions) {
	message.Attachments = Attachments{
		Attachment{
			Text:    attachmentMessage,
			Actions: actions,
		},
	}

}

// CreateButton Create a button for Slack
func CreateButton(name, buttonText, buttonType, buttonValue string) *Action {

	return &Action{
		Name:  name,
		Text:  buttonText,
		Type:  buttonType,
		Value: buttonValue,
	}

}

// SendMessage Sends POST request to Slack
func SendMessage(message *MessageStruct) {

	// Create JSON payload
	payload, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	fmt.Println(bytes.NewBuffer(payload))

	// Send Payload as byte stream to specified SLACK URL
	_, postErr := http.Post(os.Getenv("SLACK_URL"), "application/json", bytes.NewBuffer(payload))
	if postErr != nil {
		panic(postErr)
	}
}

// AddLineBreak Add line break to string
func AddLineBreak() string {
	return "\n"
}

// AddSpacer Adds double break to string
func AddSpacer() string {
	return "\n\n"
}
