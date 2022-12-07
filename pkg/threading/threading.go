package threading

import "sync"

type Threading struct {
	wg sync.WaitGroup
	// Channel
	ch chan interface{}
	// Number of threads
	ThreadMax int
}

func InitThreading(ThreadMax int) *Threading {
	if ThreadMax <= 0 {
		println("ThreadMax must be greater than 0")
		ThreadMax = 1
	} else if ThreadMax > 100 {
		print("ThreadMax is too large, it will be set to 100\n")
		ThreadMax = 100
	}
	return &Threading{
		ThreadMax: ThreadMax,
		ch:        make(chan interface{}, ThreadMax),
		wg:        sync.WaitGroup{},
	}

}

func (t *Threading) Add() {
	t.wg.Add(1)
	t.ch <- "ok"
}
func (t *Threading) Done() {
	t.wg.Done()
	<-t.ch
}
func (t *Threading) Wait() {
	t.wg.Wait()
}
func (t *Threading) GetChannel() chan interface{} {
	return t.ch
}
func (t *Threading) GetThreadMax() int {
	return t.ThreadMax
}
func (t *Threading) SetThreadMax(ThreadMax int) {
	t.ThreadMax = ThreadMax
}
