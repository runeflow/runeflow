package hostname

import (
	"os"

	"github.com/runeflow/runeflow/message"
)

// Monitor is a hostname monitor
type Monitor struct{}

// NewMonitor creates a new hostname monitor
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Sample checks the hostname
func (m *Monitor) Sample(stats *message.StatsPayload) {
	hn, err := os.Hostname()
	if err != nil {
		return
	}
	stats.Hostname = &hn
}
