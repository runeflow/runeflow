package cpu

import (
	"log"

	"github.com/runeflow/runeflow/message"
)

// Monitor is an implementation of the Monitor interface which keeps track of
// CPU usage
type Monitor struct {
	lastMeasurement *cpuMeasurement
}

// NewMonitor creates a new CPU monitor
func NewMonitor() *Monitor {
	meas, err := measureCPU()
	if err != nil {
		log.Printf("error measuring CPU usage: %v", err)
		return &Monitor{}
	}
	return &Monitor{lastMeasurement: meas}
}

// Sample computes the percent of CPU non-idle time since the last sample was
// taken
func (mon *Monitor) Sample(stats *message.StatsPayload) {
	meas, err := measureCPU()
	if err != nil {
		log.Printf("error measuring CPU usage: %v", err)
		return
	}
	stats.CPU = &message.CPUStats{}
	if mon.lastMeasurement != nil {
		delta := meas.total - mon.lastMeasurement.total
		if delta >= 0 {
			nonIdle := float64(delta - meas.idle + mon.lastMeasurement.idle)
			stats.CPU.Used = nonIdle / float64(delta)
		}
	}
	mon.lastMeasurement = meas
}
