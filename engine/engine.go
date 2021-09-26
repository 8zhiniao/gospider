package engine

import (
	"github.com/8zhiniao/gospider/fetcher"
	"log"
)

func Run(seed... Request){

	var requests []Request
	for _,r := range seed{
		requests = append(requests,r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching %s",r.Url)

		//调用fetch去获取网页内容
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetcher:error"+"fetcher.url %s,%v",r.Url,err)
			continue
		}

		//把requests加入到队列里面
		parserResult := r.ParserFunc(body)
		requests = append(requests,parserResult.Requests...)

		//对应的结果Items循环出来
		for _,item := range parserResult.Items{
			log.Printf("go get item %v",item)
		}
	}


}
