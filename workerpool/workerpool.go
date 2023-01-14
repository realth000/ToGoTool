package workerpool

import (
	"sync"
)

type Work = func()

type payloadPool chan Work

type WorkerPool struct {
	Size int
	pool payloadPool
	wg   sync.WaitGroup
}

// NewWorkerPool returns a new WorkerPool, pool size is size.
func NewWorkerPool(size int) *WorkerPool {
	return &WorkerPool{
		Size: size,
		pool: make(payloadPool, size),
	}
}

// Add adds a work to WorkerPool, asynchronously.
func (w *WorkerPool) Add(work Work) {
	w.wg.Add(1)
	// Use wg to wait for goroutine finish.
	go w.add(work)
}

// AddSync adds a work to WorkerPool, synchronously.
func (w *WorkerPool) AddSync(work Work) {
	w.wg.Add(1)
	w.add(work)
}

func (w *WorkerPool) add(work Work) {
	defer w.wg.Done()
	w.pool <- work
	work()
	<-w.pool
}

// Wait waits for all tasks in workerpool to finish.
func (w *WorkerPool) Wait() {
	w.wg.Wait()
}
