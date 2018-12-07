package client

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/runeflow/runeflow/message"
)

var (
	// ErrConnectionNotOpen indicates an attempt was made to use the client's
	// connection while it was not open.
	ErrConnectionNotOpen = errors.New("connection not open")

	// ErrAlreadyConnected indicates that a connection has already been
	// established
	ErrAlreadyConnected = errors.New("already connected")
)

// A Client handles communication with the server over a websocket
type Client struct {
	apiKey   string
	endpoint string
	conn     *websocket.Conn
}

// NewClient creates a new client with the provided API key
func NewClient(apiKey, endpoint string) *Client {
	return &Client{apiKey: apiKey, endpoint: endpoint}
}

// Connect establishes the websocket connection
func (c *Client) Connect() error {
	if c.conn != nil {
		return ErrAlreadyConnected
	}
	authHeader := []string{fmt.Sprintf("token %s", c.apiKey)}
	conn, _, err := websocket.DefaultDialer.Dial(c.endpoint, http.Header{
		"Authorization": authHeader,
	})
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

// CloseConn closes the websocket connection
func (c *Client) CloseConn() error {
	if c.conn == nil {
		return ErrConnectionNotOpen
	}
	err := c.conn.Close()
	c.conn = nil
	return err
}

// SendClose sends a message requesting that the server terminate the connection
func (c *Client) SendClose() error {
	if c.conn == nil {
		return ErrConnectionNotOpen
	}
	msg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	err := c.conn.WriteMessage(websocket.CloseMessage, msg)
	if err != nil {
		return fmt.Errorf("error sending close message: %v", err)
	}
	return nil
}

// ReadMessage reads an incoming message
func (c *Client) ReadMessage() (*message.Message, error) {
	if c.conn == nil {
		return nil, ErrConnectionNotOpen
	}
	var msg message.Message
	if err := c.conn.ReadJSON(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

// SendStats sends some statistics
func (c *Client) SendStats(stats map[string]interface{}) error {
	if c.conn == nil {
		return ErrConnectionNotOpen
	}
	statsMessage, err := message.NewStatsMessage(&message.StatsPayload{
		Stats:     stats,
		Timestamp: time.Now(),
	})
	if err != nil {
		return err
	}
	return c.conn.WriteJSON(statsMessage)
}

// SendCommandAck tells the server we received the command
func (c *Client) SendCommandAck(id string) error {
	mes, err := message.NewCmdAckMessage(id)
	if err != nil {
		return err
	}
	return c.conn.WriteJSON(mes)
}

// SendCommandResult sends the result of the command to the server
func (c *Client) SendCommandResult(r *message.CmdResultPayload) error {
	msg, err := message.NewCmdResultMessage(r)
	if err != nil {
		return err
	}
	return c.conn.WriteJSON(msg)
}
