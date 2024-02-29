package queue

import (
	"errors"
	"sync"
)

var ErrOutOfRange = errors.New("index out of range")

type Queue struct {
	mu    sync.RWMutex
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) <= 0 || q.items == nil {
		return 0
	}
	return len(q.items)
}

func (q *Queue) Enqueue(values ...interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, values...)
}

func (q *Queue) Dequeue() (interface{}, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	if len(q.items) <= 0 {
		return nil, ErrOutOfRange
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) Peek() (interface{}, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	if len(q.items) <= 0 {
		return nil, ErrOutOfRange
	}
	return q.items[0], nil
}

func (q *Queue) Contains(value interface{}) bool {
	q.mu.RLock()
	defer q.mu.RUnlock()
	for _, item := range q.items {
		if item == value {
			return true
		}
	}
	return false
}

func (q *Queue) Remove(value interface{}) error {
	q.mu.RLock()
	defer q.mu.RUnlock()
	for i, item := range q.items {
		if item == value {
			q.items = append(q.items[:i], q.items[i+1:]...)
			return nil
		}
	}
	return ErrOutOfRange
}

func (q *Queue) Purge() {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.items != nil || cap(q.items) > 0 {
		q.items = make([]interface{}, 0)
	}
}
