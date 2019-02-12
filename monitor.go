package main

const (
	ringIntervalS = 10
	ringMaxLen    = 6 * 2
)

type Monitor struct {
	// TODO expand to map section to ring
	ring *Ring
}

func NewMonitor(ringNumIntervals int) *Monitor {
	return &Monitor{
		// ring: NewRing(ringNumIntervals, ringIntervalLenS),
		// each item in ring is 10s, need to be size 1 so can perform operations on it
		ring: NewRing(ringNumIntervals),
	}
}

func (m *Monitor) ProcessLine(ts int64) {
	m.ring.AddToRing(ts)
}
