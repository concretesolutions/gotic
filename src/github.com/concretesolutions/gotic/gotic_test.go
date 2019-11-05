package gotic

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockedExecBashPipedCommand struct {
	mock.Mock
}

func TestExecBashPipedCommand(t *testing.T) {
	output, err := ExecBashPipedCommand("echo Arthur", true)

	if err != nil {
		t.Errorf("got '%v'", output)
	}

	if output != "Arthur" {
		t.Errorf("got '%v'", output)
	}
}

func TestExecBashPipedCommandIgnoreExitCode(t *testing.T) {
	output, err := ExecBashPipedCommandIgnoreExitCode("echo Arthur", true)

	if err != nil {
		t.Errorf("got '%v'", output)
	}

	if output != "Arthur" {
		t.Errorf("got '%v'", output)
	}
}

func TestPrompt(t *testing.T) {
	Prompt("question", "ls", false)

}
