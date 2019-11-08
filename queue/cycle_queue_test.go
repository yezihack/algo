package queue

import "testing"

func TestCycleQueue(t *testing.T) {
	//环形队列实际长度为5,但只能添加4个元素
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
	for i := 5; i < 9; i++ {
		if b := q.Push(i); !b {
			t.Errorf("push: expect:true, actual:%v\n", b)
		}
	}
	if b := q.Push(10); b {
		t.Errorf("push: expect:true, actual:%v\n", b)
	}
	t.Log(q.Display())
}
