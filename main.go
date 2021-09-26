package main

import (
	"github.com/8zhiniao/gospider/engine"
	"github.com/8zhiniao/gospider/zhenai/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCity,
	})
}
