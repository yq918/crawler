package parser

import (
	"engine"
	"regexp"
)
const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`


/**
解析城市
 */
func ParseCity(contents []byte) engine.ParserResult  {
	compile := regexp.MustCompile(cityRe)
	all := compile.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _,m := range all{
		name := string(m[2])
		result.Items = append(result.Items,"User: "+ name)
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				 return ParseProfile(c,name)
			},
		})
	}
	return  result
}