package apache

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type apacheStatus map[string]string

// serverStatus will attempt to make an HTTP request to the server status
// endpoint provided by mod_status and returns a map of the stats returned
func serverStatus() (apacheStatus, error) {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1/server-status?auto", nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(body), "\n")
	stats := map[string]string{}
	for _, line := range lines {
		sides := strings.Split(line, ": ")
		if len(sides) == 2 {
			stats[sides[0]] = sides[1]
		}
	}
	return stats, nil
}

func (s apacheStatus) getInt(key string) int64 {
	if val, ok := s[key]; ok {
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return intVal
		}
	}
	return 0
}

func (s apacheStatus) getFloat(key string) float64 {
	if val, ok := s[key]; ok {
		floatVal, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return floatVal
		}
	}
	return 0
}
