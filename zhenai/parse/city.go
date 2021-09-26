package parse

import (
	"github.com/8zhiniao/gospider/engine"
	"regexp"
)

const city = "<a href=\"(http://album.zhenai.com/u/[0-9]+)\"+[^>]*>([^<]+)</a>"

func ParserCityUser(content []byte) engine.PaseResult {
    res := regexp.MustCompile(city)
	matcher := res.FindAllSubmatch(content,-1)

	result := engine.PaseResult{}

	for _,m := range matcher{
		name := string(m[2])
		result.Items = append(result.Items,"user"+name)
		result.Requests = append(result.Requests,engine.Request{
			string(m[1]),
			engine.NilParser,
		})
	}

	return result
}