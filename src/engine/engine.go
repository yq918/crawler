package engine

import (
	"fetcher"
	"log"
)

func Run(seeds ...Request)  {
   var Requests  []Request
   for _,r := range seeds{
         Requests = append(Requests,r)
   }
   for len(Requests) > 0 {
         r := Requests[0]
         Requests = Requests[1:]

	    body, e := fetcher.Fetch(r.Url)
	    log.Printf("fetch URl:%s\n",r.Url)
	    if e != nil{
	    	log.Printf("error fetch URl:%s,%v",r.Url,e)
	    	continue
		}
	   parserRet := r.ParserFunc(body)
	   Requests = append(Requests,parserRet.Requests...)

	   for _,itme := range parserRet.Items{
	         log.Printf("Got Item:%v\n",itme)
	   }
   }
}
