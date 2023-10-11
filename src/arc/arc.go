// Adaptive Replacement Cache Implementation
//
// Dependencies: lru.go, utility.go
//
// Description:
// ARC (Adaptive Replacement Cache) is a fixed-size cache with even
// better performance than LRU for most work loads. It keeps track of
// both recently used (T1) and frequently used keys (T2), plus a recent
// eviction history (ghost list) B1 and B2 for each corresponding cache
// entries. It is adapative in the sense that it will dynamically prefer
// extending its cache size to accomodate for entries that populate T1 more
// than T2, or vice-versa.

package arc

import (
	"fmt"
	"os"
)

type ARC struct {
	size int // size is the fixed number of key-value pairs the cache stores

	// p is the index that starts off in the center, and shifts to accomidate more
	// entries in T1, less entries in T2, or vice-versa. This is the factor that
	// determines preferences between recently-accessed items and frequently accessed.
	p int

	t1 *LRU // T1 is the list for recently accessed items, with LRU eviction
	t2 *LRU // T2 is the list for frequently accessed items, with LRU eviction
	// ghost lists implemented as Hash Sets, represents "metadata" of cache
	b1    map[string]bool // B1 is the set of keys evicted from t1
	b2    map[string]bool // B2 is the set of keys evicted from t2
	stats *Stats          // maintains stats associated with hits/misses
}

// NewARC creates an ARC of the given size
func NewARC(size int) *ARC {

	t1 := NewLru(size) // max size of t1 or t2 is the full cache size
	t2 := NewLru(size)
	b1 := make(map[string]bool)
	b2 := make(map[string]bool)
	stats := &Stats{0, 0}

	arc := &ARC{
		size:  size,
		p:     size / 2, // favor recency/frequency equally at start
		t1:    t1,
		t2:    t2,
		b1:    b1,
		b2:    b2,
		stats: stats,
	}

	return arc
}

// Get returns the value associated with the given key, if it exists.
//
func (arc *ARC) Get(key string) ([]byte, bool) {

	valuet1, t1Contains := arc.t1.Get(key)
	valuet2, t2Contains := arc.t2.Get(key)

	// if t1Contains, we promote it to t2 (since it was accessed a 2nd time)
	if t1Contains {
		arc.t1.Remove(key)       // remove from T1
		arc.t2.Set(key, valuet1) // place in T2
		arc.stats.Hits++
		return valuet1, true

	} else if t2Contains { // if in t2, stays in t2
		arc.stats.Hits++
		return valuet2, true
	}

	arc.stats.Misses++
	// if not in either t1 or t2, then was a miss
	return nil, false
}

// Set puts a key-value pair into cache.
func (arc *ARC) Set(key string, value []byte) {

	// check for key in T1 or T2
	
	t1Contains := arc.t1.Contains(key)
	t2Contains := arc.t2.Contains(key)

	// similar to Get, move key to T2 if was in T1 and update value
	if t1Contains {
		arc.t1.Remove(key)
		arc.t2.Set(key, value)
		return
	} else if t2Contains {
		arc.t2.Set(key, value)
		return
	}

	lenB1 := len(arc.b1)
	lenB2 := len(arc.b2)

	// if missed on T1 and T2, check the ghost lists B1, B2 to update p
	if arc.b1[key] {
		// since B1 contained key, increase p to favor T1
		var increaseBy int

		if lenB2 > lenB1 {
			increaseBy = lenB2 / lenB1
		} else {
			increaseBy = 1
		}

		arc.p = min(arc.p+increaseBy, arc.size) // don't want to exceed size, so take arc.size upper bound

		// if arc len at max size, need to evict a key from T2 to increase T1
		if arc.Len() >= arc.size {
			arc.evictToGhost("B2")
		}

		// Delete from B1, and add key to T2 (since accessed 2nd time)
		delete(arc.b1, key)
		arc.t2.Set(key, value)
		return

	} else if arc.b2[key] {
		// Since B2 contained key, decrease p to favor T2
		var decreaseBy int

		if lenB2 < lenB1 {
			decreaseBy = lenB1 / lenB2
		} else {
			decreaseBy = 1
		}

		arc.p = max(arc.p-decreaseBy, 0) // Can't have negative, so take 0 as lower bound

		// need to evict a key from T1 to increase T2
		if arc.Len() >= arc.size {
			arc.evictToGhost("B1")
		}

		// Delete key from B2, move to T2 (means it was accessed min of 3 times)
		delete(arc.b2, key)
		arc.t2.Set(key, value)
		return
	}

	// Case when encountering a brand new key 

	// if size going to be exceeded, introduce new key into T1 (after evicting from B1)
	if arc.Len() >= arc.size {
		arc.evictToGhost("B1")
	}

	// Add to the recently seen list
	arc.t1.Set(key, value)
	arc.handleGhostLists() // control the size of B lists from growing indefinitely
}

// B1 and B2 are the metadata of evicted keys, to prevent size of this metadata
// growing indefinitely, we start removing keys (at random) from it
func (arc *ARC) handleGhostLists() {
	if len(arc.b1) > arc.size {
		removeRandKey(arc.b1)
	}
	if len(arc.b2) > arc.size {
		removeRandKey(arc.b2)
	}
}

// evictToGhost is used to evict a key from the
// cache (T1 + T2) into B1 or B2 (from passed in whichList)
func (arc *ARC) evictToGhost(whichList string) {
	
	if arc.t1.Len() > 0 && whichList == "B1" {
		key, ok := arc.t1.RemoveLRU()

		if ok {
			arc.b1[key] = true
		}
	} else if arc.t2.Len() > 0 && whichList == "B2" {
		key, ok := arc.t2.RemoveLRU()

		if ok {
			arc.b2[key] = true
		}
	}
}

// Len returns the number of entries in the ARC
func (arc *ARC) Len() int {
	lenT1 := arc.t1.Len()
	lenT2 := arc.t2.Len()

	return lenT1 + lenT2
}

// Remove removes and returns the value associated with the given key, if it exists.
// If key not in the cache, returns nil,false
func (arc *ARC) Remove(key string) ([]byte, bool) {

	// figure whether in T1 or T2, and remove
	val1, ok1 := arc.t1.Remove(key)
	if ok1 {
		return val1, true
	}

	val2, ok2 := arc.t2.Remove(key)
	if ok2 {
		return val2, true
	}

	// if not in either t1 or t2, then was a miss
	return nil, false
}

// returns to the size of the ARC cache
func (arc *ARC) MaxSize() int {
	return arc.size
}

// Stats returns statistics about how many search hits and misses have occurred.
func (arc *ARC) Stats() *Stats {
	return arc.stats
}

// report hits/misses from Get calls to stdout
func (arc *ARC) ReportStats() {
	fmt.Println("ARC Hits/Misses")
	fmt.Println("Number of Hits:", arc.stats.Hits)
	fmt.Println("Number of Misses:", arc.stats.Misses)
	fmt.Println("Percentage of Hits:", 100 * float64(arc.stats.Hits) / float64(arc.stats.Misses + arc.stats.Hits))
}

// for debugging
func (arc *ARC) invariant() bool {

	// check duplicate keys between T1 and T2
	t1L := arc.t1.ReturnKeys()
	t2L := arc.t2.ReturnKeys()

	keepTrackKeys := make(map[string]bool)
	for _, k := range t1L {
		keepTrackKeys[k] = true
	}

	for _, key := range t2L {
		if keepTrackKeys[key] { // means intersection
			fmt.Fprintf(os.Stderr, "A key in T1 found in T2")
			return false
		}
	}

	// fmt.Println("T1 size:", arc.t1.Len())
	// fmt.Println("T2 size:", arc.t2.Len())
	// fmt.Println("B1 size:", len(arc.b1))
	// fmt.Println("B1:", arc.b1)

	// fmt.Println("B2 size:", len(arc.b2))
	// fmt.Println("B2:", arc.b2)

	// fmt.Println("size:", arc.size)

	return true
}
