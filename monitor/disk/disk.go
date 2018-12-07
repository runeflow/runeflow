package disk

import (
	"io/ioutil"
	"log"
	"strings"
	"syscall"

	"github.com/runeflow/runeflow/message"
)

// A Monitor is simply a Monitor implementation that checks for disk usage
type Monitor struct{}

// NewMonitor creates a new disk monitor
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Sample retrieves the current disk stats
func (d *Monitor) Sample(m *message.StatsPayload) {
	disks, err := getMountedDisks()
	if err != nil {
		log.Printf("error getting mounts: %v", err)
		return
	}
	stableDisks := filterStable(disks)
	for _, d := range stableDisks {
		err := statfs(d)
		if err != nil {
			log.Printf("stat error: %v", err)
			return
		}
	}
	m.Disk = stableDisks
}

func getMountedDisks() ([]*message.DiskStats, error) {
	data, err := ioutil.ReadFile("/proc/mounts")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	disks := []*message.DiskStats{}
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 3 {
			disks = append(disks, &message.DiskStats{
				Mountpoint: fields[1],
				Filesystem: fields[2],
			})
		}
	}
	return disks, nil
}

func statfs(d *message.DiskStats) error {
	var stat syscall.Statfs_t
	err := syscall.Statfs(d.Mountpoint, &stat)
	if err != nil {
		return err
	}
	d.Blocks = int64(stat.Blocks)
	d.BlockSize = int64(stat.Bsize)
	d.BlocksFree = int64(stat.Bfree)
	return nil
}

func filterStable(disks []*message.DiskStats) []*message.DiskStats {
	stableDisks := []*message.DiskStats{}
	for _, disk := range disks {
		if isStable(disk.Filesystem) {
			stableDisks = append(stableDisks, disk)
		}
	}
	return stableDisks
}
