package arc

// necessary utility functions used in ARC

// use stats to keep track of hits and misses (same from Assignment 3)
type Stats struct {
	Hits   int
	Misses int
}

func (stats *Stats) Equals(other *Stats) bool {
	if stats == nil && other == nil {
		return true
	}
	if stats == nil || other == nil {
		return false
	}
	return stats.Hits == other.Hits && stats.Misses == other.Misses
}

// gets max of two integers
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// gets min of two integers
func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

// remove a random key from a set
func removeRandKey(b map[string]bool) {
	// map iteration is random, so first key in iterable is deleted
	for k, _ := range b {
		delete(b, k)
		return
	}
}
