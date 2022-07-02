package parsers

import (
	"go_code/project11/crawler/engine"
	"regexp"
	"strings"
)

const cityList = `{"linkContent":"([^"]+)","linkURL":"(http:\\u002F\\u002Fwww.zhenai.com\\u002Fzhenghun\\u002F[a-zA-Z0-9]+)"}`

// 解析请求
func ParseRequest(contents []byte) engine.ParseResult {
	index := 0
	re := regexp.MustCompile(cityList)
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, match := range all {
		index++
		if index == 3 {
			break
		}
		// 城市添加到item中
		result.Item = append(result.Item, string(match[1]))
		// 每个城市的url构造成新的requestA
		result.Requests = append(result.Requests, engine.Request{
			Url:       strings.ReplaceAll(string(match[2]), "\\u002F", "/"),
			ParseFunc: ParseCity,
		})
	}
	return result
}
