package main

import (
	"bufio"
	"fmt"
	"github.com/8zhiniao/gospider/fetcher"
	"github.com/8zhiniao/gospider/zhenai/parse"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"regexp"

	//"golang.org/x/text"
	//"golang.org/x/net/html"
)


func main(){
    url := "http://www.zhenai.com/zhenghun"
	all, err := fetcher.Fetch(url)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n",all)
	e := parse.ParseCity(all)
	fmt.Printf("%v\n",e.Requests)
	fmt.Printf("%v\n",e.Items)
	fmt.Printf("%v",e)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
    return e
}

/*
解析城市列表
*/
func printCityAll(content []byte){
	//sstring := "<a target=\"_blank\" href=\"http://www.zhenai.com/zhenghun/chongqing\" data-v-f53df81a>重庆</a>"
	compareString :="<a href=\"http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+\" [^>]*>[^<]+</a>"
	//<a href="http://www.zhenai.com/zhenghun/weihai" data-v-1573aa7c>威海</a>
	re := regexp.MustCompile(compareString)
	allString := re.FindAll(content,-1)
	for _,m := range allString{
		fmt.Println(string(m))
	}
	fmt.Println(len(allString))
}

func printSubCityAll(content []byte){
	//sstring := "<a target=\"_blank\" href=\"http://www.zhenai.com/zhenghun/chongqing\" data-v-f53df81a>重庆</a>"
	compareString :="<a href=\"(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)\" [^>]*>([^<]+)</a>"
	//<a href="http://www.zhenai.com/zhenghun/weihai" data-v-1573aa7c>威海</a>
	re := regexp.MustCompile(compareString)
	allString := re.FindAllSubmatch(content,-1)
	for _,m := range allString{
		fmt.Printf("%s %s\n",m[2],m[1])
	}
	fmt.Println(len(allString))
}
