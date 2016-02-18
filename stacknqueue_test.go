package stacknqueue_test
import (
	"testing"
	sq "github.com/morganhein/stacknqueue"
	"sync"
)

func TestQueueNotThreadsafe(t *testing.T) {
	q := sq.NewStackNQueue(false)
	for i := 0; i < 200; i++ {
		q.Queue(i)
	}
	if q.Len() != 200 {
		t.Error("Wrong length count detected. Expected", 200, "got", q.Len())
	}
	//Test popping from the front
	for i := 0; i < 200; i++ {
		data := q.Pop()
		if (data != i) {
			t.Error("Wrong data value detected. Expected ",
				i,
				" but got ",
				data)
		}
	}
	if q.Len() != 0 {
		t.Error("Wrong length count detected. Expected", 200, "got", q.Len())
	}
	//Make sure the StackNQueue is empty
	for i := 0; i < 200; i++ {
		q.Queue(i)
	}
	for i := 199; i > -1; i-- {
		data := q.Dequeue()
		if (data != i) {
			t.Error("Wrong data value detected. Expected ",
				i,
				" but got ",
				data)
		}
	}
	if q.Len() != 0 {
		t.Error("Wrong length count detected. Expected", 200, "got", q.Len())
	}
}

func TestStackNotThreadsafe(t *testing.T) {
	q := sq.NewStackNQueue(false)
	for i := 0; i < 200; i++ {
		q.Push(i)
	}
	for i := 199; i > -1; i-- {
		data := q.Pop()
		if data != i {
			t.Error("Wrong data value detected. Expected ", i, "got", data)
		}
		// Test peek functionality
		if q.Len() != 0 && q.Peek() != (i - 1) {
			t.Error("Wrong next data value detected. Expected ",
				i - 1,
				" but got ",
				data)
		}
	}
	if q.Len() != 0 {
		t.Error("Wrong length count detected. Expected", 200, "got", q.Len())
	}
}

func TestHelpers(t *testing.T) {
	q := sq.NewStackNQueue(false)
	for i := 0; i < 200; i++ {
		q.Push(i)
	}
	if q.IsEmpty() == true {
		t.Error("IsEmpty returned incorrect result. Expected to be false but got true.")
	}
	q.Empty()
	if q.IsEmpty() != true {
		t.Error("IsEmpty returned incorrect result. Expected to be true but got false.")
	}
}

func TestQueueThreadsafe(t *testing.T) {
	q := sq.NewStackNQueue(true)
	var wg sync.WaitGroup

	wg.Add(3)

	go fillList(q, 200, &wg)
	go fillList(q, 200, &wg)
	go fillList(q, 200, &wg)

	wg.Wait()

	if q.Len() != 600 {
		t.Error("Wrong length count detected. Expected", 600, "got", q.Len())
	}

	wg.Add(2)
	c := make(chan bool)
	go emptyList(q, &wg, c)

	fillList(q, 1500, &wg)

	c <- true

	wg.Wait()

	if q.Len() != 0 {
		t.Error("Wrong length count detected. Expected", 0, "got", q.Len())
	}
}

func fillList(q *sq.StackNQueue, c int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < c; i++ {
		q.Queue(i)
	}
}

func emptyList(q *sq.StackNQueue, wg *sync.WaitGroup, c chan bool) {
	defer wg.Done()
	select {
	case m := <-c:
		if m == true {
			emptyHelper(q)
			close(c)
			return
		}
	default:
		emptyHelper(q)
	}
}

func emptyHelper(q *sq.StackNQueue) {
	for q.IsEmpty() == false {
		_ = q.Pop()
	}
}
