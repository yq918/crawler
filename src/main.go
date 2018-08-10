package main

import (
	"engine"
	"zhenai/parser"
	"scheduler"
)

const URL = "http://www.zhenai.com/zhenghun"

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
     //  Url:URL,
     //  ParserFunc:parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		   Scheduler:&scheduler.SimpleScheduler{},
		   WorkCount:10,
	}
	e.Run(engine.Request{
		Url:URL,
		ParserFunc:parser.ParseCityList,
	})

}


