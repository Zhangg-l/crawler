package engine

import (
	"fmt"
	"go_code/project11/crawler/fetcher"
	"go_code/project11/crawler/model"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, v := range seeds {
		requests = append(requests, v)
	}
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		// 解析页面
		content, err := fetcher.Fetch(req.Url)
		if err != nil {
			log.Println("Fetch err :", err)
			continue
		}
		// 提取url
		parseResult := req.ParseFunc(content)
		if pro, ok := parseResult.Item[0].(model.Profile); ok {
			fmt.Printf("%#v", pro)
		} else {
			requests = append(requests, parseResult.Requests...)
			fmt.Println(parseResult.Item)
		}

		// for _, v := range parseResult.Item {
		// 	log.Printf("city:%v\n", v)
		// }
	}

}
