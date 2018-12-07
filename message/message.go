package message

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// StatsMessage messages are sent from the client to the server. It contains
	// information about the current machine state.
	StatsMessage = "stats"

	// CmdMessage messages are sent from the server to the client. It specifies
	// an action to be carried out by the client, such as rebooting the machine.
	CmdMessage = "cmd"

	// CmdAckMessage messages are sent from the client to the server to indicate
	// that a command has been received.
	CmdAckMessage = "cmd-ack"

	// CmdResultMessage messages are sent from the client to the server after an
	// action has been carried out indicating the result.
	CmdResultMessage = "cmd-result"
)

// A Message is the top-level structure used to communicate between server and
// client. It has a type field and a payload, with the payload structure being
// dependent on the message type.
type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// A CmdPayload represents the payload for a MsgTypeCmd message. It contains a
// server-generated unique ID and the type of action to be carried out.
type CmdPayload struct {
	ID     string `json:"id"`
	Action string `json:"action"`
}

// A CmdAckPayload represents the payload for a MsgTypeCmdAck message. It
// simply consists of the unique ID of the action being acknowledged from the
// MsgTypeCmd message.
//
// The client must always send an ack message before it carries out the
// associated command. If an ack is not received by the server after an
// interval, the server may resend the cmd with the same ID. The client must
// execute received commands with the same ID at most once.
type CmdAckPayload struct {
	ID string `json:"id"`
}

// A CmdResultPayload represents the payload for a MsgTypeCmdResult message. It
// contains the ID of the command of which the result is being reported and
// whether the command was successful.
//
// It is possible that this message will not be sent, e.g. for a reboot
// command, the client might send a failure message but if the reboot is
// successful, the connection will be terminated.
type CmdResultPayload struct {
	ID        string `json:"id"`
	Error     string `json:"error"`
	StdOutput string `json:"stdOutput"`
	StdError  string `json:"stdError"`
}

// A StatsPayload represents the payload for a StatsMessage message. It
// contains a timestamp and a dictionary of statistics being reported.
type StatsPayload struct {
	Stats     map[string]interface{} `json:"stats"`
	Timestamp time.Time              `json:"timestamp"`
}

// ParseCmd parses the message payload as CmdPayload
func (m *Message) ParseCmd() (*CmdPayload, error) {
	if m.Type != CmdMessage {
		return nil, fmt.Errorf("can't decode %s payload as %s", m.Type, CmdMessage)
	}
	p := CmdPayload{}
	if err := json.Unmarshal(m.Payload, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// ParseStats parses the message payload as StatsPayload
func (m *Message) ParseStats() (*StatsPayload, error) {
	if m.Type != StatsMessage {
		return nil, fmt.Errorf("can't decode %s payload as %s", m.Type, StatsMessage)
	}
	p := StatsPayload{}
	if err := json.Unmarshal(m.Payload, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// NewStatsMessage creates a new stats message with the provided payload
func NewStatsMessage(p *StatsPayload) (*Message, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return &Message{
		Type:    StatsMessage,
		Payload: body,
	}, nil
}

// NewCmdAckMessage creates a new ack message for the command id received
func NewCmdAckMessage(id string) (*Message, error) {
	ack := CmdAckPayload{
		ID: id,
	}
	body, err := json.Marshal(ack)
	if err != nil {
		return nil, err
	}
	return &Message{
		Type:    CmdAckMessage,
		Payload: body,
	}, nil
}

// NewCmdResultMessage creates a new result message for the command
func NewCmdResultMessage(r *CmdResultPayload) (*Message, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return &Message{
		Type:    CmdResultMessage,
		Payload: body,
	}, nil
}
