package scheduler

import "engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(e engine.Request) {
	go func() {s.workerChan <- e}()
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(c chan engine.Request)  {
      s.workerChan = c
}


