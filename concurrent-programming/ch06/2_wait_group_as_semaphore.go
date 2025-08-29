package main

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func New(n int) *Semaphore {
	return &Semaphore{
		permits: n,
		cond: sync.NewCond(&sync.Mutex{}), 
	}
}

func (s *Semaphore) Acquire() {
	s.cond.L.Lock()
	for s.permits <= 0 {
		s.cond.Wait()
	}
	s.permits--
	s.cond.L.Unlock()	
}

func (s *Semaphore) Release() {
	s.cond.L.Lock()
	s.permits++
	s.cond.Signal()
	s.cond.L.Unlock()
}

type WaitGroup struct {
	semaphore *Semaphore
}

func NewWaitGroup(size int) *WaitGroup {
	return &WaitGroup{semaphore: New(1 - size)}
}

func (wg *WaitGroup) Wait() {
	wg.semaphore.Acquire()
}

func (wg *WaitGroup) Done() {
	wg.semaphore.Release()
}

func main() {
	wg := NewWaitGroup(4) 

	for i := 1; i <= 4; i++ {
		go _doWork(i, wg)
	}
	wg.Wait()
	fmt.Println("All works are completed")
}

func _doWork(id int, wg *WaitGroup) {
	fmt.Println(id, "Done")
	wg.Done()
}