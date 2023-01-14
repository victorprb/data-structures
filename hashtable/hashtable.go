package hashtable

type HashTable struct {
	buckets []bucket
	size    int
}

func New(size int) *HashTable {
	return &HashTable{
		buckets: make([]bucket, size, size),
		size:    size,
	}
}

func (h *HashTable) Add(key string, value any) {
	index := h.hash(key)

	h.buckets[index].add(key, value)
}

func (h *HashTable) Get(key string) any {
	index := h.hash(key)

	entry := h.buckets[index].get(key)
	if entry != nil {
		return entry.value
	}

	return nil
}

func (h *HashTable) Delete(key string) {
	index := h.hash(key)

	h.buckets[index].delete(key)
}

func (h *HashTable) hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % h.size
}

type bucket struct {
	head *bucketEntry
}

type bucketEntry struct {
	key   string
	value any
	next  *bucketEntry
}

func (b *bucket) add(key string, value any) {
	entry := b.get(key)
	if entry != nil {
		entry.value = value
		return
	}

	newEntry := &bucketEntry{
		key:   key,
		value: value,
		next:  b.head,
	}

	b.head = newEntry
}

func (b *bucket) get(key string) *bucketEntry {
	current := b.head

	for current != nil {
		if current.key == key {
			return current
		}

		current = current.next
	}

	return nil
}

func (b *bucket) delete(key string) {

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previous := b.head

	for previous.next != nil {
		if previous.next.key == key {
			previous.next = previous.next.next
			return
		}

		previous = previous.next
	}
}
