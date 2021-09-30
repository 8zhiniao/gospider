package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main(){

	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless",false),
		chromedp.Flag("ignore-certificate-errors",true),
		chromedp.WindowSize(1920,1080),
		chromedp.Flag("blink-settings","imagesEnabled=false"),
	)

	allocCtx, cancel1 := chromedp.NewExecAllocator(context.TODO(), options...)
	defer cancel1()

	ctx, cancel2 := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel2()

	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.cnblogs.com/"),
		chromedp.WaitVisible(`#footer`, chromedp.ByID),
		chromedp.Nodes(`a.post-item-title`, &nodes),

	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get nodes:",len(nodes))
	//print titles
	for _,node := range nodes{
		fmt.Println(node.Children[0].NodeValue,":",node.AttributeValue("href"))
		fmt.Println("----------------------------------------")
	}

	var res string
	err2 := chromedp.Run(ctx,setcookies("https://www.cnblogs.com/",&res,"cookie1", "value1"))
	if err2 != nil {
		panic(err2)
	}


}

func setcookies(host string, res *string, cookies ...string) chromedp.Tasks {
	if len(cookies)%2 != 0 {
		panic("length of cookies must be divisible by 2")
	}
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// add cookies to chrome
			for i := 0; i < len(cookies); i += 2 {
				err := network.SetCookie(cookies[i], cookies[i+1]).
					WithExpires(&expr).
					WithDomain("localhost").
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		// navigate to site
		chromedp.Navigate(host),
		// read the returned values
		chromedp.Text(`#result`, res, chromedp.ByID, chromedp.NodeVisible),
		// read network values
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetAllCookies().Do(ctx)
			if err != nil {
				return err
			}

			for i, cookie := range cookies {
				log.Printf("chrome cookie %d: %+v", i, cookie)
			}

			return nil
		}),
	}
}
