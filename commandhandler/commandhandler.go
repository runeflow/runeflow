package commandhandler

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/runeflow/runeflow/command"
	"github.com/runeflow/runeflow/message"
)

type handleFunc func() error

type commandState struct {
	action         string
	err            string
	stdOut         string
	stdErr         string
	receivedAt     time.Time
	acknowledgedAt time.Time
	completedAt    time.Time
	reportedAt     time.Time
}

// A CommandHandler handles commands and keeps track of their states
type CommandHandler struct {
	commands map[string]*commandState
}

// NewCommandHandler creates a new command handler
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		commands: map[string]*commandState{},
	}
}

// HandleCommand handles a command
func (h *CommandHandler) HandleCommand(id, action string) *message.CmdResultPayload {
	log.Printf("Handle Command %v\n", action)
	if _, ok := h.commands[id]; ok {
		return &message.CmdResultPayload{
			ID:    id,
			Error: fmt.Sprintf("command %s has already been handled", id),
		}
	}
	var stdOut, stdErr string
	var err error
	switch action {
	case command.Reboot:
		stdOut, stdErr, err = reboot()
	case command.RestartApache:
		stdOut, stdErr, err = restartApache()
	case command.RestartMySQL:
		stdOut, stdErr, err = restartMySQL()
	}
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	h.commands[id] = &commandState{
		stdOut:     stdOut,
		stdErr:     stdErr,
		err:        errMsg,
		action:     action,
		receivedAt: time.Now().UTC(),
	}
	return &message.CmdResultPayload{
		ID:     id,
		Error:  errMsg,
		StdOut: stdOut,
		StdErr: stdErr,
	}
}

func run(command string, args ...string) (string, string, error) {
	cmd := exec.Command(command, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err := cmd.Run()
	if err != nil {
		return "", "", err
	}
	return stdout.String(), stderr.String(), nil
}

func reboot() (string, string, error) {
	return run("shutdown", "-r", "now")
}

func restartApache() (stdout, stderr string, err error) {
	return run("service", "apache2", "restart")
}

func restartMySQL() (stdout, stderr string, err error) {
	return run("service", "myswl", "restart")
}
