//
// See https://pkg.go.dev/container/heap
package priorityqueue

import (
	"container/heap"
	"sync"
	"time"

	"github.com/ppcamp/go-workers/priorityqueue/intqueue"
)

type Job interface {
	Process()
}

type Worker interface {
	Start()
	Push(job Job)
	Shutdown()
	Consume()
}

type worker struct {
	jobs       chan Job
	close      chan bool
	finishTask chan bool
	ticker     time.Ticker
	queue      intqueue.IntHeap
	N          int
	wg         *sync.WaitGroup
}

func NewWorker(workers, jobs int, delay time.Duration) Worker {
	wg := &sync.WaitGroup{}
	wg.Add(workers)

	w := &worker{
		jobs:       make(chan Job, jobs),
		close:      make(chan bool, 1),
		finishTask: make(chan bool, 1),
		queue:      intqueue.IntHeap{},
		N:          workers,
		ticker:     *time.NewTicker(delay),
		wg:         wg,
	}

	heap.Init(&w.queue)
	return w
}

func (w *worker) Shutdown() {
	w.close <- true
	w.wg.Wait()
}
func (w *worker) Push(job Job) { heap.Push(&w.queue, job) }

func (w *worker) Consume() {
	for {
		select {
		case <-w.close:
			w.wg.Done()
			return
		default:
			for job := range w.jobs {
				job.Process()
			}
		}
	}
}

func (w *worker) Produce() {
	for {
		select {
		case <-w.close:
			return
		case <-w.ticker.C:
			w.jobs <- w.queue.Pop().(Job)
		case <-w.finishTask:
			w.jobs <- w.queue.Pop().(Job)
		}
	}
}

func (w *worker) Start() {
	for i := 0; i < w.N; i++ {
		go w.Consume()
	}

	go w.Produce()
}
