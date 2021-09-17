package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		fmt.Println("err")
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s/n",all)
	//fmt.Println(all)
}
