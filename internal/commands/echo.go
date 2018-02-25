package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/types"
	"github.com/nlopes/slack"
	"log"
)

type Echo struct {
	Event  types.Event
	Input  types.Input
	Client *slack.Client
}

func (command Echo) Run() {
	message := fmt.Sprintf("<@%s>: %s", command.Event.User, command.Input.Args)

	postParams := slack.PostMessageParameters{}
	_, _, err := command.Client.PostMessage(command.Event.Channel, message, postParams)

	if err != nil {
		log.Fatal(err)
	}
}
