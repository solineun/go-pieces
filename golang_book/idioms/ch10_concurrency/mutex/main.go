package main

import "sync"

func main() {
	
}

func scoreBoardManager(in <-chan func(map[string]int), done chan struct{}) {
	scoreboard := map[string]int{}
	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(scoreboard)
		}
	}
}

type ChannelScoreBoardManager chan func(map[string]int)

func NewChannelScoreBoardManager() (ChannelScoreBoardManager, func()) {
	ch := make(ChannelScoreBoardManager)
	done := make(chan struct{})
	go scoreBoardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func (csm ChannelScoreBoardManager) Update(name string, val int) {
	csm <- func (m map[string] int)  {
		m[name] = val
	}
}

func (csm ChannelScoreBoardManager) Read(name string) (int, bool) {
	var val int
	var ok bool
	done := make(chan struct{})
	csm <- func (m map[string]int)  {
		val, ok = m[name]
		close(done)
	}
	<-done
	return val, ok
}

type MutexScoreBoardManager struct {
	l sync.RWMutex
	scoreboard map[string]int
}

func NewMutexScoreBoardManager() *MutexScoreBoardManager {
	return &MutexScoreBoardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *MutexScoreBoardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutexScoreBoardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}