package hw03frequencyanalysis

type Heap struct {
	size      int
	capacity  int
	entrances []*Entrance
}

func (h *Heap) Left(i int) int {
	return 2*i + 1
}

func (h *Heap) Right(i int) int {
	return 2*i + 2
}

func (h *Heap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Insert(e *Entrance) {
	h.size++
	i := h.size - 1
	h.entrances[i] = e

	for i != 0 && h.Compare(i, h.Parent(i)) {
		h.entrances[h.Parent(i)], h.entrances[i] = h.entrances[i], h.entrances[h.Parent(i)]
		i = h.Parent(i)
	}
}

func (h *Heap) Heapify(i int) {
	left := h.Left(i)
	right := h.Right(i)

	smallest := i
	if left < h.size && h.Compare(left, smallest) {
		smallest = left
	}

	if right < h.size && h.Compare(right, smallest) {
		smallest = right
	}

	if smallest != i {
		h.entrances[i], h.entrances[smallest] = h.entrances[smallest], h.entrances[i]
		h.Heapify(smallest)
	}
}

func (h *Heap) Extract() *Entrance {
	if h.size <= 0 {
		return nil
	}

	returnValue := h.entrances[0]
	h.entrances[0] = h.entrances[h.size-1]
	h.entrances[h.size-1] = nil
	h.size--

	h.Heapify(0)

	return returnValue
}

func (h *Heap) Compare(i, j int) bool {
	return h.entrances[i].Count > h.entrances[j].Count ||
		(h.entrances[i].Count == h.entrances[j].Count &&
			h.entrances[i].Word < h.entrances[j].Word)
}

func ConstructHeap(capacity int) *Heap {
	return &Heap{
		size:      0,
		capacity:  capacity,
		entrances: make([]*Entrance, capacity),
	}
}
