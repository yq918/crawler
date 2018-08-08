package parser

import (
	"engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/\w+)[^>]*">([^<]+)</a>`

func ParseCityList(contents []byte ) engine.ParserResult {
	compile := regexp.MustCompile(cityListRe)
	all := compile.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	i := 0
	for _,m := range all{
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParserFunc:ParseCity,
		})
		i++
		if i >= 6{
			break
		}
	}

   return  result
}
