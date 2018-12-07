package memory

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/runeflow/runeflow/message"
)

// Monitor is an implementation of the Monitor interface
type Monitor struct{}

// NewMonitor creates a new memory monitor
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Sample collects memory usage
func (m *Monitor) Sample(stats *message.StatsPayload) {
	contents, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		log.Printf("error reading meminfo: %v", err)
		return
	}
	lines := strings.Split(string(contents), "\n")
	stats.Memory = &message.MemoryStats{}
	for _, line := range lines {
		readInto(&stats.Memory.MemTotal, "MemTotal", line)
		readInto(&stats.Memory.MemFree, "MemFree", line)
		readInto(&stats.Memory.SwapTotal, "SwapTotal", line)
		readInto(&stats.Memory.SwapFree, "SwapFree", line)
	}
}

func readInto(dst *int64, prefix, line string) {
	if strings.HasPrefix(line, prefix) {
		value, err := extractMemoryValue(line)
		if err != nil {
			return
		}
		*dst = value
	}
}

func extractMemoryValue(line string) (int64, error) {
	sides := strings.Split(line, ":")
	if len(sides) != 2 {
		return 0, errors.New("unrecognized line format")
	}
	rhs := strings.TrimSpace(sides[1])
	rhs = strings.Split(rhs, " ")[0]
	i, err := strconv.ParseInt(rhs, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse value: %v", err)
	}
	return i, nil
}
