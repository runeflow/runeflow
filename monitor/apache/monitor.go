package apache

import (
	"log"
)

// Monitor is a monitor for the Apache web server
type Monitor struct{}

// Stats are the apache statistics we retrieve
type Stats struct {
	IsRunning         bool    `json:"isRunning"`
	Uptime            int64   `json:"uptime"`
	RequestsPerSecond float64 `json:"requestsPerSecond"`
}

// NewMonitor creates a new apache monitor
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Sample collects apache statistics
func (m *Monitor) Sample() interface{} {
	stats := &Stats{
		IsRunning: isRunning(),
	}
	status, err := serverStatus()
	if err != nil {
		log.Printf("server status error: %v", err)
		return nil
	}
	stats.Uptime = status.getInt("ServerUptimeSeconds")
	stats.RequestsPerSecond = status.getFloat("ReqPerSec")
	return stats
}
