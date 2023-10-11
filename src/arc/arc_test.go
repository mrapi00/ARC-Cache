/******************************************************************************
 * lru_test.go
 * Author:
 * Usage:    `go test`  or  `go test -v`
 * Description:
 *    An incomplete unit testing suite for lru.go. You are welcome to change
 *    anything in this file however you would like. You are strongly encouraged
 *    to create additional tests for your implementation, as the ones provided
 *    here are extremely basic, and intended only to demonstrate how to test
 *    your program.
 ******************************************************************************/

package arc

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

/******************************************************************************/
/*                                Functions                                   */
/******************************************************************************/

// Fails test t with an error message if fifo.MaxStorage() is not equal to capacity
// func checkCapacity(t *testing.T, cache Cache, capacity int) {
// 	max := cache.MaxStorage()
// 	if max != capacity {
// 		t.Errorf("Expected %s to have %d MaxStorage, but it had %d", cacheType(cache), capacity, max)
// 	}
// }
/******************************************************************************/
/*                                  Tests                                     */
/******************************************************************************/

// function for testing LRU cache algorithm
func TestLRU(t *testing.T) {
	fmt.Println("Test LRU\n--------------")
	capacity := 10
	lru := NewLru(10)

	if lru.MaxSize() != capacity {
		t.Errorf("WRONG CAPACITY")
	}

}

// function for increasing probabiliy of getting same key
func mapToSame(val int) int {
	offset := 20 - val
	if offset < 0 {
		offset *= -1
	}

	if offset < 9 {
		return 20
	}

	return val
}

// function for testing ARC vs. LRU comparison (control)
func TestARCControl(t *testing.T) {
	fmt.Println("Test ARC vs LRU\n-----------------")
	arc := NewARC(20)
	lru := NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	lru.ReportStats()
	arc.ReportStats()
	if lru.Stats().Hits < arc.Stats().Hits {
		fmt.Println("ARC outperforms LRU (good)")
	} else {
		fmt.Println("LRU outperforms ARC (idk)")
	}

}

// function for testing ARC vs. LRU with different Set-Get ratios
func TestARCSetGetRatio(t *testing.T) {
	fmt.Println("Test ARC vs LRU Set-Get Ratio\n-----------------")
	arc := NewARC(20)
	lru := NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Set-Get Ratio 1:2", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Set-Get Ratio 1:3", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Set-Get Ratio 1:4", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

		rand.Seed(time.Now().UnixNano())
		key2 = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Set-Get Ratio 1:5", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)
	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Set-Get Ratio 5:1", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)
	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Set-Get Ratio 4:1", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)
	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Set-Get Ratio 3:1", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key = fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Set(key, []byte(""))
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)
	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Set-Get Ratio 2:1", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)
	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Set-Get Ratio 1:1", LRUHitRate(lru), ARCHitRate(arc))
}

// function for testing ARC vs. LRU with different cache sizes
func TestARCCacheSize(t *testing.T) {
	fmt.Println("Test ARC vs LRU Cache Size\n-----------------\nControl")
	// MAIN CHANGE IS HERE: CHANGED FROM 20 to 30
	arc := NewARC(20)
	lru := NewLru(20)
	// ALSO NEED TO CHANGE KEY SPACE TO 200
	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(300))) // k30
		arc.Set(key, []byte(""))                          // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(300)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Cache Size: 20", LRUHitRate(lru), ARCHitRate(arc))

	if lru.Stats().Hits < arc.Stats().Hits {
		fmt.Println("ARC outperforms LRU (good)\n-----------------")
	} else {
		fmt.Println("LRU outperforms ARC (idk)")
	}

	arc = NewARC(30)
	lru = NewLru(30)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(300))) // k30
		arc.Set(key, []byte(""))                          // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(300)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Cache Size: 30", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(60)
	lru = NewLru(60)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(300))) // k30
		arc.Set(key, []byte(""))                          // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(300)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Cache Size: 60", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(120)
	lru = NewLru(120)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(300))) // k30
		arc.Set(key, []byte(""))                          // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(300)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Cache Size: 120", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(200)
	lru = NewLru(200)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(300))) // k30
		arc.Set(key, []byte(""))                          // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(300)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Cache Size: 200", LRUHitRate(lru), ARCHitRate(arc))

}

// function for testing ARC vs. LRU with different range of keys
func TestARCRangeKeys(t *testing.T) {
	fmt.Println("Test ARC vs LRU Range of Keys\n-----------------")
	arc := NewARC(20)
	lru := NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		// MAIN CHANGE IS HERE: changed from 40 to 80
		key := fmt.Sprint("k", mapToSame(rand.Intn(80))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(80)))

		arc.Get(key2)
		lru.Get(key2)

	}
	fmt.Printf("%v,%v,%v\n", "Range of Keys: 80", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(160))) // k30
		arc.Set(key, []byte(""))                          // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(160)))

		arc.Get(key2)
		lru.Get(key2)

	}
	fmt.Printf("%v,%v,%v\n", "Range of Keys: 160", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(320))) // k30
		arc.Set(key, []byte(""))                          // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(320)))

		arc.Get(key2)
		lru.Get(key2)

	}
	fmt.Printf("%v,%v,%v\n", "Range of Keys: 320", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(400))) // k30
		arc.Set(key, []byte(""))                          // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(400)))

		arc.Get(key2)
		lru.Get(key2)

	}
	fmt.Printf("%v,%v,%v\n", "Range of Keys: 400", LRUHitRate(lru), ARCHitRate(arc))

}

// function for testing ARC vs. LRU with different number of requests
func TestARCMoreIterations(t *testing.T) {
	fmt.Println("Test ARC vs LRU Iterations\n-----------------")
	arc := NewARC(20)
	lru := NewLru(20)

	// MAIN CHANGE IS HERE (changed from 5k to 10k)
	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Iterations: 10000", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 15000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Iterations: 15000", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 20000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Iterations: 20000", LRUHitRate(lru), ARCHitRate(arc))

	arc = NewARC(20)
	lru = NewLru(20)

	for i := 0; i < 25000; i++ {
		rand.Seed(time.Now().UnixNano())
		key := fmt.Sprint("k", mapToSame(rand.Intn(40))) // k30
		arc.Set(key, []byte(""))                         // value: ""
		lru.Set(key, []byte(""))

		rand.Seed(time.Now().UnixNano())
		key2 := fmt.Sprint("k", mapToSame(rand.Intn(40)))

		arc.Get(key2)
		lru.Get(key2)

	}

	if !arc.invariant() {
		t.Errorf("INVARIANT VIOLATED")
	}

	fmt.Printf("%v,%v,%v\n", "Iterations: 25000", LRUHitRate(lru), ARCHitRate(arc))
}

func LRUHitRate(lru *LRU) float64 {
	hits := float64(lru.stats.Hits)
	misses := float64(lru.stats.Misses)
	return hits / (hits + misses) * 100
}

func ARCHitRate(arc *ARC) float64 {
	hits := float64(arc.stats.Hits)
	misses := float64(arc.stats.Misses)
	return hits / (hits + misses) * 100
}

// function for testing ARC vs. LRU with Wikipedia 2019 trace
func TestWikipediaTrace(t *testing.T) {
	fmt.Println("Testing first 10m lines of trace with different cache sizes")
	cacheSize := 500
	batch := 2
	for cacheSize <= 5_000_000 {
		lru, arc, err := testOnTrace("wiki2019.tr", cacheSize, batch*10_000_000, (batch+1)*10_000_000)
		if err != nil {
			fmt.Println("Encountered error: ", err)
		}

		fmt.Printf("%v,%v,%v\n", cacheSize, LRUHitRate(lru), ARCHitRate(arc))

		cacheSize *= 10
	}
}

func testOnTrace(filename string, size int, start, end int) (*LRU, *ARC, error) {
	arc := NewARC(size)
	lru := NewLru(size)

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Create a new buffered reader
	reader := bufio.NewReader(file)

	i := 0

	// Read the rows from the file
	for i < end {
		i++
		// Read a line from the file
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, nil, err
		}

		if i < start {
			continue
		}

		// Split the line into columns
		columns := strings.Split(line, " ")

		// Convert the columns to integers

		key := columns[1]

		if _, arcHit := arc.Get(key); !arcHit {
			arc.Set(key, []byte{})
		}
		if _, lruHit := lru.Get(key); !lruHit {
			lru.Set(key, []byte{})
		}

	}

	return lru, arc, nil
}
