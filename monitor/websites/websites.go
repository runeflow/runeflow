package websites

import (
	"log"
	"os/exec"
	"strings"

	"github.com/runeflow/runeflow/message"
)

// Monitor keeps track of websites that we can discover as being hosted on this machine
type Monitor struct{}

// NewMonitor creates a new website monitor
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Sample attempts to discover websites being hosted on this machine
func (m *Monitor) Sample(stats *message.StatsPayload) {
	apacheHosts, err := discoverApacheVHosts()
	if err != nil {
		log.Printf("error discovering apache vhosts: %v", err)
		return
	}
	stats.Websites = apacheHosts
}

func discoverApacheVHosts() ([]string, error) {
	out, err := exec.Command("apache2ctl", "-t", "-D", "DUMP_VHOSTS").CombinedOutput()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(out), "\n")
	vhostsStarted := false
	discovered := map[string]struct{}{}
	for _, line := range lines {
		if vhostsStarted {
			hostFields := strings.Fields(line)
			if len(hostFields) >= 2 {
				discovered[hostFields[1]] = struct{}{}
			}
		}
		if strings.HasPrefix(line, "VirtualHost configuration") {
			vhostsStarted = true
		}
	}
	uniqueDiscovered := []string{}
	for k := range discovered {
		uniqueDiscovered = append(uniqueDiscovered, k)
	}
	return uniqueDiscovered, nil
}
