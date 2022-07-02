package main

import (
	"go_code/project11/crawler/engine"
	"go_code/project11/crawler/zhenai/parsers"
)

// 单机版
func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parsers.ParseRequest,
	})

}
