package queue_test

import (
	"sync"
	"testing"

	"github.com/aliffatulmf/go-queue"
)

func TestQueueSequence(t *testing.T) {
	t.Run("Sequential Integers", func(t *testing.T) {
		q := queue.NewQueue()

		for i := 0; i < 10; i++ {
			q.Enqueue(i)
		}

		for i := 0; i < 10; i++ {
			got, err := q.Dequeue()
			if err != nil {
				t.Errorf("got: %v, want: %v", err, nil)
			}
			if got != i {
				t.Errorf("got: %v, want: %v", got, i)
			}
		}
	})

	t.Run("Sequential Random Type", func(t *testing.T) {
		q := queue.NewQueue()

		args := []interface{}{0, "1", 2, "3", 4, "5", 6, "7", 8, "9"}
		q.Enqueue(args...)

		for _, i := range args {
			got, err := q.Dequeue()
			if err != nil {
				t.Errorf("got: %v, want: %v", err, nil)
			}
			if got != i {
				t.Errorf("got: %v, want: %v", got, i)
			}
		}
	})
}

func TestQueueInGoroutine(t *testing.T) {
	q := queue.NewQueue()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			q.Enqueue(i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			q.Dequeue()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
