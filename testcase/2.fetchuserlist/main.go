package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	//"golang.org/x/text"
	//"golang.org/x/net/html"
)


func main(){

	url := "https://album.zhenai.com/u/1446944988"
	client := &http.Client{}
	request, err2 := http.NewRequest("GET", url, nil)
	if err2 != nil {
		panic(err2)
	}
    client.Do(request)

	url1 := "https://album.zhenai.com/u/1446944988"
	//url2 := "http://www.baidu.com"
	resp, err := http.Get(url1)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		fmt.Println("err")
	}
	//resp.Body
	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		panic(err)
	}
	fmt.Printf("%s/n",all)
	//decodeRune, _ := utf8.DecodeRune(all)

	//fmt.Println(decodeRune)
	//fmt.Println(string(all))

	//re := regexp.MustCompile("<a target=\"_blank\" href=\"http://www.zhenai.com/zhenghun/chongqing\" data-v-f53df81a>重庆</a> ")
	//allString := re.FindAllString(string(all),-1)
	//fmt.Println(allString)
	//printSubCityAll(all)
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
