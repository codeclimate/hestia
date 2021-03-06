package commands

import (
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"testing"
)

func TestEcho(t *testing.T) {
	notifier := notifiers.Test{}

	command := Echo{User: "test", Input: types.Input{Args: "hello"}, Notifier: &notifier}
	command.Run()

	messages := notifier.Messages
	expected := "hello"

	if messages[0] != expected {
		t.Fatalf("Expected `%s`, but received `%s`", expected, messages[0])
	}
}
