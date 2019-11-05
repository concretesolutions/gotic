package gotic

import (
	"testing"
)

func TestExecBashPipedCommand(t *testing.T) {
	output, err := ExecBashPipedCommand("ls", false)

	if err != nil {
		t.Errorf("got '%t'", output)
	}

	if output == nil {
		t.Errorf("got '%t'", output)
	}
}

func TestExecBashPipedCommandIgnoreExitCode(t *testing.T) {
	output, err := ExecBashPipedCommandIgnoreExitCode("ls", false)

	if err != nil {
		t.Errorf("got '%t'", output)
	}

	if output == nil {
		t.Errorf("got '%t'", output)
	}
}

func TestPrompt(t *testing.T) {
	output, err := Prompt("question", "ls", false)

	if err != nil {
		t.Errorf("got '%t'", output)
	}

	if output == nil {
		t.Errorf("got '%t'", output)
	}
}
