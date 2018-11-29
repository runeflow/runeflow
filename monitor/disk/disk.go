package disk

import (
	"io/ioutil"
	"log"
	"strings"
	"syscall"
)

// A Monitor is simply a Monitor implementation that checks for disk usage
type Monitor struct{}

// NewMonitor creates a new disk monitor
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Sample retrieves the current disk stats
func (d *Monitor) Sample() interface{} {
	disks, err := getMountedDisks()
	if err != nil {
		log.Printf("error getting mounts: %v", err)
		return nil
	}
	stableDisks := filterStable(disks)
	for _, d := range stableDisks {
		err := d.statfs()
		if err != nil {
			log.Printf("stat error: %v", err)
			return nil
		}
	}
	return stableDisks
}

func getMountedDisks() ([]*Disk, error) {
	data, err := ioutil.ReadFile("/proc/mounts")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	disks := []*Disk{}
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 3 {
			disks = append(disks, &Disk{
				Mountpoint: fields[1],
				Filesystem: fields[2],
			})
		}
	}
	return disks, nil
}

func (d *Disk) statfs() error {
	var stat syscall.Statfs_t
	err := syscall.Statfs(d.Mountpoint, &stat)
	if err != nil {
		return err
	}
	d.Blocks = stat.Blocks
	d.BlockSize = stat.Bsize
	d.BlocksFree = stat.Bfree
	return nil
}

func filterStable(disks []*Disk) []*Disk {
	stableDisks := []*Disk{}
	for _, disk := range disks {
		if isStable(disk.Filesystem) {
			stableDisks = append(stableDisks, disk)
		}
	}
	return stableDisks
}
