package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan  Request)
} 

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigMasterWorkerChan(in)

	for i:=0; i<e.WorkCount;i++{
        createWorker(in,out)
	}

	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _,item := range result.Items{
            log.Printf("Got item %v:",item)
		}
		for _,request := range result.Requests{
             e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParserResult)  {
	 go func() {
	 	for{
	 		request := <- in
			result, e := Worker(request)
			if e != nil{
				continue
			}
			out <- result
		}
	 }()
}


