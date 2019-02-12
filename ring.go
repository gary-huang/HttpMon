package main

import "log"

type Ring struct {
	cap           int
	size          int
	headIdx       int
	headTs        int64
	trafficCounts []int
}

func NewRing(capacity int) *Ring {
	return &Ring{
		cap:           capacity,
		size:          0,
		headIdx:       -1,
		headTs:        -1,
		trafficCounts: make([]int, capacity),
	}
}

func (r *Ring) isOlderThanTwoMin(curTs int64) bool {
	// check if curTs, current timestamp is older than 2 mins
	return true
}

// TODO get time conversions right
func (r *Ring) AddToRing(ts int64) {
	if r.headTs == -1 {
		r.headTs = ts
		r.headIdx = 0
	}

	// TODO change the 1 to the real tail index
	tailTs := r.headTs + int64(r.size*10)

	// evict old entries
	// TODO get the actual index to stop at, i.e. it might wrap
	lstop := r.headIdx + r.size*10
	for i := r.headIdx; i < lstop; i++ {
		if r.isOlderThanTwoMin(ts) {
			r.headIdx += 1
			r.size -= 1
			r.headTs = r.headTs + 10
		}
	}

	if ts > tailTs {
		// TODO handle wrapping/overflow
		r.size += int((ts - tailTs) / 10)
		r.trafficCounts[r.size+r.headIdx-1] = 1
	} else {
		r.trafficCounts[r.headIdx+r.size*10] = 1
	}

	// update this bucket
	idxToUpdate := (ts - r.headTs) / 10
	// TODO check for loop around to head of array
	r.trafficCounts[idxToUpdate] += 1
	log.Printf("Current Ring: %v", r)
}
