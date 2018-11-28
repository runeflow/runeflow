package cpu

import "log"

// Monitor is an implementation of the Monitor interface which keeps track of
// CPU usage
type Monitor struct {
	lastMeasurement *cpuMeasurement
}

// Stat is the result of the CPU statistics
type Stat struct {
	Used float64 `json:"used"`
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
func (mon *Monitor) Sample() interface{} {
	meas, err := measureCPU()
	if err != nil {
		log.Printf("error measuring CPU usage: %v", err)
		return nil
	}
	stat := &Stat{}
	if mon.lastMeasurement != nil {
		delta := meas.total - mon.lastMeasurement.total
		if delta >= 0 {
			nonIdle := float64(delta - meas.idle + mon.lastMeasurement.idle)
			stat.Used = nonIdle / float64(delta)
		}
	}
	mon.lastMeasurement = meas
	return stat
}
