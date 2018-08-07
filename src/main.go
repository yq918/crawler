package main

import (
	"engine"
	"zhenai/parser"
)

const URL = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
       Url:URL,
       ParserFunc:parser.ParseCityList,
	})
}
