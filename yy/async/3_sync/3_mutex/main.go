package main

import (
	"sync"
	"time"
)

type storage struct {
	m  map[int]int
	mu sync.Mutex
}

func (s *storage) write(i int) {
	s.m[0] = i
}

func (s *storage) check() {
	if s.m[0] < 0 {

	}
}

func main() {
	s := storage{
		m: make(map[int]int),
	}

	go func() {
		for i := 0; i < 1e6; i++ {
			s.mu.Lock()
			defer s.mu.Unlock()
			s.write(i)
		}
	}()

	go func() {
		for i := 0; i < 1e6; i++ {
			s.mu.Lock()
			defer s.mu.Unlock()
			s.check()
		}
	}()

	time.Sleep(time.Second / 2)
}