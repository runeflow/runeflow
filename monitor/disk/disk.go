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

// A Disk represents the stats for a mounted filesystem
type Disk struct {
	Mountpoint string `json:"mountpoint"`
	Filesystem string `json:"filesystem"`
	Blocks     int64  `json:"blocks"`
	BlockSize  int64  `json:"blockSize"`
	BlocksFree int64  `json:"blocksFree"`
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
	d.Blocks = int64(stat.Blocks)
	d.BlockSize = int64(stat.Bsize)
	d.BlocksFree = int64(stat.Bfree)
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
