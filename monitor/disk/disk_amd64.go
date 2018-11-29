package disk

// A Disk represents the stats for a mounted filesystem
type Disk struct {
	Mountpoint string `json:"mountpoint"`
	Filesystem string `json:"filesystem"`
	Blocks     uint64 `json:"blocks"`
	BlockSize  int64  `json:"blockSize"`
	BlocksFree uint64 `json:"blocksFree"`
}
