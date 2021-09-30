package main

import (
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"log"
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


}