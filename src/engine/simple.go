package engine

import (
	"fetcher"
	"log"
)

type SimpleEngine struct {

}

func (s SimpleEngine) Run(seeds ...Request)  {
   var Requests  []Request
   for _,r := range seeds{
         Requests = append(Requests,r)
   }
   for len(Requests) > 0 {
         r := Requests[0]
         Requests = Requests[1:]

         parserRet, e := Worker(r)
	     if e != nil{
			   continue
		   }
		 Requests = append(Requests,parserRet.Requests...)

		  for _,itme := range parserRet.Items{
				 log.Printf("Got Item:%v\n",itme)
		   }
   }
}


func Worker(r Request) (ParserResult,error){
	log.Printf("fetch URl:%s\n",r.Url)
	body, e := fetcher.Fetch(r.Url)
	if e != nil{
		return ParserResult{},e
	}
	parserRet := r.ParserFunc(body)
	return parserRet,nil
}