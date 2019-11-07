package gotic

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
Gotic package

Provides a proper interface to execute shell command lines
*/

// BuildCommand builds up a exec.Cmd
var (
	BuildCommand = buildCommand
	logFatal     = log.Fatal
)

func buildCommand(command string) (cmd *exec.Cmd) {
	return exec.Command("bash", "-c", command)
}

// ExecBashPipedCommand executes a simple command or a piped bash command
func ExecBashPipedCommand(command string, showOutput bool) (string, error) {
	cmd := BuildCommand(command)

	out, err := cmd.CombinedOutput()

	count := strings.Count(string(out), "\n")

	if count == 1 {
		out = []byte(strings.ReplaceAll(string(out), "\n", ""))
	}

	if out != nil && showOutput {
		log.Printf("Done executing command: %s", command)
		log.Printf("Output:%s", string(out))
	}

	if err != nil {
		s := fmt.Sprintf("ExecBashPipedCommand failed with %s", err)
		logFatal(s)
	}

	return string(out), err
}

// ExecBashPipedCommandIgnoreExitCode executes a simple command or a piped bash command
func ExecBashPipedCommandIgnoreExitCode(command string, showOutput bool) (string, error) {
	cmd := BuildCommand(command)

	out, err := cmd.CombinedOutput()

	count := strings.Count(string(out), "\n")

	if count == 1 {
		out = []byte(strings.ReplaceAll(string(out), "\n", ""))
	}

	if out != nil && showOutput {
		log.Printf("Done executing command: %s", command)
		log.Printf("Output:%s", string(out))
	}

	if err != nil && showOutput {
		log.Printf("ExecBashPipedCommand failed with %v\n", err)
	}

	return string(out), err
}

// ExecShellScript executes a shell script file
func ExecShellScript(shFilePath string, showOutput bool) (string, string) {
	cmd := BuildCommand(shFilePath)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("ExecShellScript failed with %s\n", err)
	}

	if out != nil {
		fmt.Printf("Done executing shellscript file %s\n", shFilePath)
		fmt.Printf("Out \n%s\n", string(out))
	}
	return string(out), fmt.Sprintf("%s", err)
}

// Prompt asks to execute a commands
func Prompt(question, command string, showOutput bool) {
	fmt.Printf("==> %s '%s'? (y/n)\n", question, command)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	if strings.Compare("Y", text) == 0 || strings.Compare("y", text) == 0 {
		_, _ = ExecBashPipedCommand(command, showOutput)
		log.Printf("Command '%s' finished\n", command)
	}
}
