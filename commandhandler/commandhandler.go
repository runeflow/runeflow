package commandhandler

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/runeflow/runeflow/command"
)

type handleFunc func() error

type commandState struct {
	action         string
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
func (h *CommandHandler) HandleCommand(id, action string) error {
	if _, ok := h.commands[id]; ok {
		return fmt.Errorf("command %s has already been handled", id)
	}
	switch action {
	case command.Reboot:
	}
	h.commands[id] = &commandState{
		action:     action,
		receivedAt: time.Now(),
	}
	return nil
}

func reboot() error {
	//exec.Cmd("shutdown", "-r", "now")
	return exec.Command("echo", "hello").Run()
}
