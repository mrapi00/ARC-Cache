// Adapted from cos316 Assignment 3

// In this version of LRU, the limit of the LRU cache is
// determined by the number of key-value pair entries as opposed to 
// a precize number of bytes (since the ARC algorithm that uses LRU
// allocates more memory based on the number of key-value pair entries).
//
// 

package arc

import "fmt"

// LRU is a fixed-size in-memory cache with last recently used eviction
type LRU struct {
	size    int // number of entries (k,v pairs) in LRU that can stored
	sentinel *Node // a sentinel node, sentinel.next = first node, sentinel.prev = last node
	mapNode  map[string]*Node // maps key to node holding the value
	stats    *Stats // maintains stats associated with hits/misses
}

// helper node class, doubly linked
type Node struct {
	prev  *Node
	next  *Node
	key   string
	value []byte
}

// NewLRU returns a pointer to a new LRU with a capacity to store limit bytes
func NewLru(size int) *LRU {
	lru := &LRU{
		size,
		&Node{},
		make(map[string]*Node),
		&Stats{0, 0},
	}
	lru.sentinel.prev = lru.sentinel
	lru.sentinel.next = lru.sentinel
	return lru
}

// MaxSize returns the number of entries supported by LRU
func (lru *LRU) MaxSize() int {
	return lru.size
}

// if a key is accessed, update it in the linked list as the most recently used key
func (lru *LRU) updateMRU(node *Node) {
		lru.detachNode(node)
		node.prev = lru.sentinel.prev
		node.next = lru.sentinel
		lru.sentinel.prev.next = node
		lru.sentinel.prev = node
}

// Get returns the value associated with the given key, if it exists.
// ok is true if a value was found and false otherwise.
func (lru *LRU) Get(key string) (value []byte, contains bool) {
	// search key, value pair by map
	node, contains := lru.mapNode[key]
	if contains {
		value = node.value
		lru.stats.Hits += 1
		lru.updateMRU(node)
	} else {
		lru.stats.Misses += 1
	}

	return
}

// Contain just checks if key is in the cache. Doesn't update recency or hits/misses
func (lru *LRU) Contains(key string) (contains bool) {
	// search key, value pair by map
	_, contains = lru.mapNode[key]
	return
}

// Remove removes and returns the value associated with the given key, if it exists.
// ok is true if a value was found and false otherwise
func (lru *LRU) Remove(key string) (value []byte, ok bool) {
	node, ok := lru.mapNode[key]
	if ok {
		// if found, remove the node, update the linked list and update all fields of the lru struct
		value = node.value
		lru.removeNode(node)
	}
	return
}


// Set associates the given value with the given key, possibly evicting values
// to make room. Returns true if the binding was added successfully, else false.
func (lru *LRU) Set(key string, value []byte) bool {
	existingNode, contains := lru.mapNode[key]
	
	needToRemove := lru.Len() == lru.size

	// if key doesn't exist, add it and update struct
	if !contains {
		// delete head (LRU) if we could only add the new key after removing the head
		if needToRemove {
			lru.deleteHead()
		}

		newNode := &Node{
			lru.sentinel.prev,
			lru.sentinel,
			key,
			value,
		}
		lru.sentinel.prev.next = newNode
		lru.sentinel.prev = newNode

		lru.mapNode[key] = newNode
	} else {
		existingNode.value = value
		lru.updateMRU(existingNode) // to move to end
	}

	return true
}

// Len returns the number of entries in the LRU.
func (lru *LRU) Len() int {
	return len(lru.mapNode)
}

// Returns the keys stored by the the LRU cache. Used for 
// debugging by the client.  
func (lru *LRU) ReturnKeys() []string{
	ans := make([]string, 0)
	for k, _ := range lru.mapNode {
		ans = append(ans, k)
	}
	// fmt.Println(ans)
	return ans
}

// Exposes the RemoveLRU function to client so they can delete the least-recently 
// used key even if size has not been exceeded
func (lru *LRU) RemoveLRU() (key string, ok bool) {
	node := lru.sentinel.next
	if node == lru.sentinel {
		return "", false
	}
	lru.detachNode(node)
	delete(lru.mapNode, node.key)
	return node.key, true
}


// Stats returns statistics about how many search hits and misses have occurred.
func (lru *LRU) Stats() *Stats {
	return lru.stats
}

func (lru *LRU) ReportStats() {
	fmt.Println("LRU Hits/Misses")
	fmt.Println("Number of Hits:", lru.stats.Hits)
	fmt.Println("Number of Misses:", lru.stats.Misses)
	fmt.Println("Percentage of Hits:", 100 * float64(lru.stats.Hits) / float64(lru.stats.Misses + lru.stats.Hits))
}

// internal helper function for displaying contents of cache when debugging
func (lru *LRU) debug() {
	currNode := lru.sentinel.next
	fmt.Print("Length: ", lru.Len(),". First Pointer",  lru.sentinel.next, ", LAST POINTER: ", lru.sentinel.prev, "\n")
	for currNode != lru.sentinel {
		fmt.Print("AT THIS NODE with key ")
		fmt.Println(currNode.key, ":", currNode.value)
		currNode = currNode.next
	}
}

/**********************************************************************************/ 
// Helper functions for List/Node manipulation

// removes a node from the linked list while leaving its key and value intact
func (lru *LRU) detachNode(node *Node) {
	if node == lru.sentinel {
		return
	}
	nextNode := node.next
	prevNode := node.prev
	nextNode.prev = prevNode
	prevNode.next = nextNode
}

// removes a node and its key and value from its containing list
func (lru *LRU) removeNode(node *Node) {
	if node == lru.sentinel {
		return
	}
	lru.detachNode(node)
	delete(lru.mapNode, node.key)
}

// helper function to deleting head of LRU, updating linked list and relevant fields of struct
func (lru *LRU) deleteHead() {
	lru.removeNode(lru.sentinel.next)
}

/**********************************************************************************/ 

