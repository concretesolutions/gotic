package gotic

import (
	"errors"
	"io"
	"log"
	"testing"
)

func TestExecBashPipedCommand(t *testing.T) {
	logFatal = log.Print
	output, err := ExecBashPipedCommand("echo Arthur", true)

	if err != nil {
		t.Errorf("got '%v'", output)
	}

	if output != "Arthur" {
		t.Errorf("got '%v'", output)
	}
}

func buildCommandMockFail(cc string) ([]byte, error) {
	return []byte("Arthur\n"), errors.New("error")
}

func buildCommandMock(cc string) ([]byte, error) {
	return []byte("Arthur\n"), nil
}

func TestExecBashPipedCommandLogFatal(t *testing.T) {
	BuildCommand = buildCommandMockFail
	logFatal = log.Print
	output, err := ExecBashPipedCommand("echo Arthur\n", true)

	if err == nil {
		t.Errorf("got '%v'", output)
	}

	if output != "Arthur" {
		t.Errorf("got '%v'", output)
	}
}

func TestExecBashPipedCommandIgnoreExitCode(t *testing.T) {
	BuildCommand = buildCommandMock
	output, err := ExecBashPipedCommandIgnoreExitCode("echo Arthur\n", true)

	if err != nil {
		t.Errorf("got '%v'", output)
	}

	if output != "Arthur" {
		t.Errorf("got '%v'", output)
	}
}

func TestExecBashPipedCommandIgnoreExitCodeLogPrintf(t *testing.T) {
	BuildCommand = buildCommandMockFail
	output, err := ExecBashPipedCommandIgnoreExitCode("echo Arthur\n", true)

	if err == nil {
		t.Errorf("got '%v'", output)
	}

	if output != "Arthur" {
		t.Errorf("got '%v'", output)
	}
}

func TestExecShellScript(t *testing.T) {
	BuildCommand = buildCommandMockFail
	out, err := ExecShellScript("echo Arthur", true)

	if err == "" {
		t.Errorf("got '%v'", out)
	}

	if out != "Arthur\n" {
		t.Errorf("got '%v'", out)
	}

}

func readPromptMock(stdin io.Reader) string {
	return "y"
}

func readPromptMockFail(stdin io.Reader) string {
	return "n"
}

func TestPrompt(t *testing.T) {
	ReadPrompt = readPromptMock
	result := Prompt("question", "ls", false)

	if result == false {
		t.Errorf("got '%v'", result)
	}
}

func TestPromptFail(t *testing.T) {
	ReadPrompt = readPromptMockFail
	result := Prompt("question", "ls", false)

	if result == true {
		t.Errorf("got '%v'", result)
	}
}
