package hostname

import "os"

// Monitor is a hostname monitor
type Monitor struct{}

// NewMonitor creates a new hostname monitor
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Sample checks the hostname
func (m *Monitor) Sample() interface{} {
	hn, err := os.Hostname()
	if err != nil {
		return nil
	}
	return hn
}
