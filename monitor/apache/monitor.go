package apache

import (
	"log"

	"github.com/runeflow/runeflow/message"
)

// Monitor is a monitor for the Apache web server
type Monitor struct{}

// NewMonitor creates a new apache monitor
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Sample collects apache statistics
func (m *Monitor) Sample(stats *message.StatsPayload) {
	stats.Apache = &message.ApacheStats{
		IsRunning: isRunning(),
	}
	status, err := serverStatus()
	if err != nil {
		log.Printf("server status error: %v", err)
		return
	}
	stats.Apache.Uptime = status.getInt("ServerUptimeSeconds")
	stats.Apache.RequestsPerSecond = status.getFloat("ReqPerSec")
}
