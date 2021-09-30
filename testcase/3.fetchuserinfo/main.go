package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func main(){

	client := &http.Client{Timeout: 10 * time.Second}
    url1 := "https://album.zhenai.com/u/1907456476"
	//url2 := "http://www.baidu.com"
	request, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		fmt.Println(err)
	}
	cookie1 := "FSSBBIl1UgzbN7NO=5iK63ZNS.IWVTFxkvCMUBlwAwBo6nLIcA9JAkM7Omtdai7i0RbwURgq.pu_o9OZXJKsjYjR4_FejWZsI0CJ3CCq; sid=c3cde05f-3e7b-4bcf-8a5e-05ff9c4f5dfb; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1632881469; ec=zCx2CWVK-1632881469809-a74ecee642b41-1771404603; _exid=1W3tFYD0oZMRK2pGBpLfeL0i/0/1wlHRUc0/VEJ9l6ByRPwcP1i2Vz/T73E3khCFIHtzviV1LL8XhSZDLhjszw==; _efmdata=Aai7nOD0FXD6Yv8PVeKdR6NqFq3bO74gzlbSSTZRjZ56adiJhrXxD8//Lt4mt82KQXcmLF6vXieItRkoUUEkzxJg9YZkib2YrOBttTZZaRI=; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1632882069; FSSBBIl1UgzbN7NP=53YHJbbmsuRVqqqmZhTnNmGp56BrC.rfCudG1595t8GC76nVm5idoJpq8y1ldnYJ230VhSnx46Yn_P6BeiDc1iHOveXZ.frYOeJ2IeKrqaQBuHEtt0QLcs_R65kbpX4hUsAmRigK8WWZ8sfWv1puIjDKiWgX8hpWwUbFPeZCStjgRKjkWOYpwvYcsVYqNiqdSTLOHqLy_ekLzHs6bfhYfYuT7s7ojilIGACFCSLOWMnEKk2qFts4C0BWnis9HOsKeE"
	request.Header.Add("cookie",cookie1)
    request.Header.Add("Referer","https://album.zhenai.com/")
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36")
	request.Header.Add("Cookie", "xxxxxx")
	request.Header.Add("X-Requested-With", "xxxx")

	resp, err1 := client.Do(request)
	if err1 != nil {
		fmt.Println(err1)
	}
	defer resp.Body.Close()
	//all, err2 := ioutil.ReadAll(resp.Body)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	fmt.Println(resp.StatusCode)
    //fmt.Println(string(all))
	bufioReader := bufio.NewReader(resp.Body)
	encode := determineEncoding(bufioReader)
	utf8Reader := transform.NewReader(bufioReader, encode.NewDecoder())
	all, err2 := ioutil.ReadAll(utf8Reader)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(all))

	compile := regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)
	allString := compile.FindAllString(string(all), -1)
	fmt.Println(allString)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)

	//bytes, err :=ioutil.ReadAll(r) //直接读取reader后就不能再读取，所以用bufferReader
	if err != nil {
		log.Printf("Fetcher error :%v", err)
		return unicode.UTF8 //默认utf8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "text/html")
	return e
}