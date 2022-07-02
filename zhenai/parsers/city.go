package parsers

import (
	"fmt"
	"go_code/project11/crawler/engine"
	"regexp"
	"strings"
)

// 解析城市的每个人的url

// <th><a href="http://album.zhenai.com/u/1958903678" target="_blank">小扬姐</a></th>
func ParseCity(contents []byte) engine.ParseResult {
	var (
		result = engine.ParseResult{}
	)
	// 解析个人列表
	userRe := regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a></th>`)
	submMatch := userRe.FindAllSubmatch(contents, -1)
	// 解析下一页
	nextPageRe := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+/[\d]+)">(下一页)</a>`)
	nextPageMatch := nextPageRe.FindAllSubmatch(contents, -1)
	fmt.Println(len(submMatch))
	for _, match := range submMatch {
		// 添加每个人的姓名
		result.Item = append(result.Item, string(match[2]))
		// 每个人的url
		result.Requests = append(result.Requests,
			engine.Request{Url: strings.ReplaceAll(string(match[1]), "\\u002F", "/"),
				ParseFunc: func(b []byte) engine.ParseResult {
					return ParseFile(b, string(match[2]))
				}})
	}

	for _, match := range nextPageMatch {
		result.Item = append(result.Item, string(match[2]))
		result.Requests = append(result.Requests,
			engine.Request{Url: strings.ReplaceAll(string(match[1]), "\\u002F", "/"),
				ParseFunc: ParseCity})
	}
	return result
}
