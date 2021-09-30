package main

import (
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"time"
)

func main()  {
	ctx,_ := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless",false))...,
		)

	ctx, _ = context.WithTimeout(ctx,30 * time.Second)
	ctx, _ = chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf))

	if err := chromedp.Run(ctx,MyTask());err != nil{
		log.Fatal(err)
		return
	}
}

func MyTask() chromedp.Tasks{

	var loginURL = "https://account.wps.cn/"
	return chromedp.Tasks{
		chromedp.Navigate(loginURL),
		chromedp.Click(`#mainWrap > div.nav_tab > span.nav_tab_item.nav_tab_main.js_nav_tab`),

	    chromedp.Click(`#wechat > span:nth-child(2)`),

	    chromedp.Click(`#dialog > div.dialog-wrapper > div > div.dialog-footer > div.dialog-footer-ok`),

	    GetCode(),
	}
}

func GetCode() chromedp.ActionFunc {

	return func(ctx context.Context) (err error) {
		var code []byte

		if err = chromedp.Screenshot(`#wximport`,&code,chromedp.ByID).Do(ctx);err != nil{
			return
		}

		if err = ioutil.WriteFile("code.png",code,0755);err != nil{
			return
		}

		return
	}
}
