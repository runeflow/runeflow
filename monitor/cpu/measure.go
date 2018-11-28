package cpu

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type cpuMeasurement struct {
	total uint64
	idle  uint64
}

func measureCPU() (*cpuMeasurement, error) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return nil, err
	}
	firstLine := strings.Split(string(contents), "\n")[0]
	fields := strings.Fields(firstLine)[1:]
	idle, err := strconv.ParseUint(fields[3], 10, 64)
	if err != nil {
		return nil, err
	}
	meas := &cpuMeasurement{idle: idle}
	for _, field := range fields {
		v, err := strconv.ParseUint(field, 10, 64)
		if err != nil {
			return nil, err
		}
		meas.total += v
	}
	return meas, nil
}
