package parsers

import (
	"go_code/project11/crawler/engine"
	"go_code/project11/crawler/model"
	"regexp"
	"strings"
)

// 从得到的文本中提取想要的数据
// 创建三个解析器
var (
	otherInfoRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div>`)
	baseInfoRe  = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
	// nameRe      = regexp.MustCompile(`<span class="nickName" data-v-4c07f04e>([^<]+)</span>`)
)

//从个人页面 解析每个人的数据
func ParseFile(contents []byte, name string) engine.ParseResult {

	profile := &model.Profile{}
	baseInfo := extractString(contents, baseInfoRe)
	otherInfo := extractString(contents, otherInfoRe)

	// nickName := string(nameRe.FindSubmatch(contents)[1])
	// if nickName != "" {
	// 	profile.Name = nickName
	// }
	// 解析base
	{
		profile.Name = name
		profile.Marriage = baseInfo[0]
		profile.Age = baseInfo[1]
		profile.Cnostellation = baseInfo[2]
		profile.Height = baseInfo[3]
		profile.Weight = baseInfo[4]
		profile.WorkPlace = baseInfo[5]
		profile.Income = baseInfo[6]
		profile.Eduction = baseInfo[8]

		profile.NativePlace = strings.Split(otherInfo[1], ":")[1]
		profile.Figure = strings.Split(otherInfo[2], ":")[1]
		profile.House = otherInfo[5]
		profile.Car = otherInfo[6]
	}
	// fmt.Println(len(baseInfo))
	// fmt.Println(len(otherInfo))
	return engine.ParseResult{Item: []interface{}{profile}}
}

func extractString(contents []byte, re *regexp.Regexp) []string {
	Strs := []string{}
	Info := re.FindAllSubmatch(contents, -1)

	for _, v := range Info {
		// fmt.Println(string(v[0]))
		Strs = append(Strs, string(v[1]))
	}
	return Strs
}
