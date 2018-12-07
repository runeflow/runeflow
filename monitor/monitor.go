package monitor

import "github.com/runeflow/runeflow/message"

// A Monitor periodically collects statistics and sends them upstream
type Monitor interface {
	Sample(*message.StatsPayload)
}
