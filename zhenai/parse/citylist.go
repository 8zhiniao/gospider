package parse

import (
	"github.com/8zhiniao/gospider/engine"
	"regexp"
)

const cityListRe = "<a href=\"(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)\" [^>]*>([^<]+)</a>"

func ParseCity(content []byte) engine.PaseResult {

	res := regexp.MustCompile(cityListRe)
	matcher := res.FindAllSubmatch(content,-1)

	result := engine.PaseResult{}

	for _,m := range matcher{
		//fmt.Printf("city:%v,url:%v:",m[2],m[1])
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParserFunc: ParserCityUser,
		})

		result.Items=append(result.Items,"city: "+string(m[2]))
	}

	return result

}
