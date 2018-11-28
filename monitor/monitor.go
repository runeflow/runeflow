package monitor

// A Monitor periodically collects statistics and sends them upstream
type Monitor interface {
	Sample() interface{}
}
