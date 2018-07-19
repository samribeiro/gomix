package unlockbug

import (
	"sync"
	"time"
)

var (
	globalMu sync.RWMutex
)

type Struct struct {
	localMu sync.RWMutex
}

func (s *Struct) CallNotify() {
	go s.Notify()
}

func (s *Struct) Notify() {
	s.localMu.Lock()
	defer s.localMu.Unlock()
	time.Sleep(time.Second * 3)
}
