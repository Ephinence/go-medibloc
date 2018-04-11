package hashheap

import "container/heap"

// Comparable is an interface that allows object to order through comparison.
type Comparable interface {
	Less(b interface{}) bool
}

// HashedHeap is a data structure for hash and priority queue.
type HashedHeap struct {
	pq   priorityQueue
	hash map[string]*entry
}

// NewHashedHeap returns HashedHeap.
func NewHashedHeap() *HashedHeap {
	return &HashedHeap{
		pq:   make(priorityQueue, 0),
		hash: make(map[string]*entry),
	}
}

// Get get value by key.
func (h *HashedHeap) Get(key string) (value Comparable) {
	e, ok := h.hash[key]
	if !ok {
		return nil
	}
	return e.value
}

// Set set value by key.
func (h *HashedHeap) Set(key string, value Comparable) {
	if _, ok := h.hash[key]; ok {
		h.Del(key)
	}

	e := &entry{
		key:   key,
		value: value,
	}
	h.hash[key] = e
	heap.Push(&h.pq, e)
}

// Del deletes value by key.
func (h *HashedHeap) Del(key string) {
	e, ok := h.hash[key]
	if !ok {
		return
	}
	heap.Remove(&h.pq, e.index)
	delete(h.hash, key)
}

// Pop pop value by priority.
func (h *HashedHeap) Pop() (value Comparable) {
	if h.pq.Len() == 0 {
		return nil
	}
	e := heap.Pop(&h.pq).(*entry)
	delete(h.hash, e.key)
	return e.value
}

// Peek peek value by priority.
func (h *HashedHeap) Peek() (value Comparable) {
	if h.pq.Len() == 0 {
		return nil
	}
	e := h.pq[0]
	return e.value
}

type entry struct {
	key   string
	value Comparable
	index int
}

type priorityQueue []*entry

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].value.Less(pq[j].value)
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	e := x.(*entry)
	e.index = len(*pq)
	*pq = append(*pq, e)
}

func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	e := (*pq)[n-1]
	e.index = -1
	*pq = (*pq)[0 : n-1]
	return e
}
