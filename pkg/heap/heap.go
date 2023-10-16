package heap

import "container/heap"

type MaxHeap []WordCount

type WordCount struct {
	Word  string
	Count int
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(WordCount))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func PopKLargestWordCounts(maxHeap *MaxHeap, k int) []WordCount {
	KLargest := make(MaxHeap, 0)
	amount := min(maxHeap.Len(), k)

	for i := 0; i < amount; i++ {
		largest := heap.Pop(maxHeap).(WordCount)
		KLargest = append(KLargest, largest)
	}

	return KLargest
}
