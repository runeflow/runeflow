package apache

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

// isRunning checks whether the apache web sever is running
func isRunning() bool {
	pid, err := ioutil.ReadFile("/var/run/apache2/apache2.pid")
	if err != nil {
		return false
	}
	out, err := exec.Command("ps", "-p", strings.TrimSpace(string(pid)), "--noheader").CombinedOutput()
	if err != nil {
		return false
	}
	return len(out) != 0
}
