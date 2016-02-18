package stacknqueue

import "sync"

type node struct {
	data interface{}
	next *node
	prev *node
}

type StackNQueue struct {
	head       *node
	tail       *node
	count      int
	threadSafe bool
	lock       *sync.Mutex
}
// NewStackNQueue creates a new LinkedList object that can be used as a
// Stack and/or a Queue and returns a pointer
func NewStackNQueue(threadSafe bool) *StackNQueue {
	q := &StackNQueue{threadSafe:threadSafe}
	if threadSafe {
		q.lock = &sync.Mutex{}
	}
	return q
}

// Len returns the number of elements in the List
func (q *StackNQueue) Len() int {
	if q.threadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()
	}

	return q.count
}

// Push inserts the value to the front of the List
func (q *StackNQueue) Push(item interface{}) {
	if q.threadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()
	}

	n := &node{data: item}

	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		q.head.prev = n
		n.next = q.head
		q.head = n
	}
	q.count++
}

// Pop removes and returns the value at the front of the List.
func (q *StackNQueue) Pop() interface{} {
	if q.threadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()

	}

	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next

	if q.head == nil {
		q.tail = nil
	}
	q.count--

	return n.data
}


// Queue adds an item to the end of the List.
// This is a FIFO action.
func (q *StackNQueue) Queue(item interface{}) {
	if q.threadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()

	}

	n := &node{data: item}

	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		n.prev = q.tail
		q.tail.next = n
		q.tail = n
	}

	q.count++
}

// Dequeue removes and returns the last item in the List.
// This is a crazy action.
func (q *StackNQueue) Dequeue() interface{} {
	if q.threadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()

	}

	//No items in the list
	if q.tail == nil {
		return nil
	}

	//Only one item in this List
	if q.head == q.tail {
		response := q.head
		q.head = nil
		q.tail = nil
		q.count--
		return response.data
	}

	response := q.tail
	q.tail = response.prev

	q.count--
	return response.data

}

// Peek returns the value at the front of the List
// without mutation. This means the value is not removed.
func (q *StackNQueue) Peek() interface{} {
	if q.threadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()

	}

	n := q.head
	if n == nil {
		return nil
	}

	return n.data
}

func (q *StackNQueue) Empty() {
	if q.threadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()

	}

	q.count = 0
	q.head = nil
	q.tail = nil
}

func (q *StackNQueue) IsEmpty() bool {
	if q.threadSafe {
		q.lock.Lock()
		defer q.lock.Unlock()
	}

	return q.count == 0
}