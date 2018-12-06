package osrelease

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var releaseFiles = []string{"/etc/os-release", "/usr/lib/os-release"}

// ReadField attempts to read the value of a field from the /etc/os-release
// file, falling back to /usr/lib/os-release if the former is not available. If
// neither file is available or the field cannot be read, an error will be
// returned.
func ReadField(field string) (string, error) {
	for _, path := range releaseFiles {
		opts, err := readReleaseFile(path)
		if err != nil {
			continue
		}
		if val, ok := opts[field]; ok {
			return val, nil
		}
		return "", fmt.Errorf("property %s was not found in %s", field, path)
	}
	return "", fmt.Errorf("no release file found")
}

func readReleaseFile(path string) (map[string]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	opts := map[string]string{}
	for _, line := range lines {
		fields := strings.SplitN(line, "=", 2)
		if len(fields) != 2 {
			continue
		}
		opts[fields[0]] = strings.TrimFunc(fields[1], func(c rune) bool {
			return c == '"'
		})
	}
	return opts, nil
}
