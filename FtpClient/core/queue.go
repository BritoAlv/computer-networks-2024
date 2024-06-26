package core

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

func (q *Queue) Enqueue(value string) {
	q.list.PushBack(value)
}

func (q *Queue) Dequeue() string {
	e := q.list.Front()
	if e != nil {
		q.list.Remove(e)
		return e.Value.(string)
	}
	return ""
}