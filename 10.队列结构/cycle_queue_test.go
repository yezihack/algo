package queue_10

import "testing"

func TestCycleQueue(t *testing.T) {
	q := NewCycleQueue(5)
	q.Push(1)
	q.Push(2)
	if q.QueueLength() != 2 {
		t.Errorf("length:expect:%d,actual:%d\n", 2, q.QueueLength())
	}
	val := q.Pop()
	if val.(int) != 1 {
		t.Errorf("pop: expect:%d, actual:%d\n", 1, val)
	}
	q.Pop()
	if q.QueueEmpty() != true {
		t.Errorf("empty: expect:true, actual:%v\n", q.QueueEmpty())
	}
	q.Push(3)

}
