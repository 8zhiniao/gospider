package main

import (
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"log"
	"time"
)

func main(){
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
    	chromedp.Flag("headless",false),
    	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chrome instance
	ctx,cancel2 := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
		)

	defer cancel2()

	//create a timeout
    ctx, cancel = context.WithTimeout(ctx,5* time.Second)
    defer cancel()

}