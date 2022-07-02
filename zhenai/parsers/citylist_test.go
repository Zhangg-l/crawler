package parsers

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCityList(t *testing.T) {
	// 获取数据
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(fmt.Sprintf("file parse fail %s", err))
	}

	// 测试解析
	result := ParseRequest(contents)
	const expectedLen = 470
	if expectedLen != len(result.Item) {
		panic(fmt.Sprintf("expected item lenth:%d;but had item %d", expectedLen, len(result.Item)))
	}
	if expectedLen != len(result.Requests) {
		panic(fmt.Sprintf("expected Requests lenth:%d;but had Requests %d", expectedLen, len(result.Item)))
	}

	// 验证获得的数据
	expectedCity := []string{"阿坝", "阿克苏", "阿拉善盟"}
	expectedUrl := []string{"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com//zhenghun//akesu", "http://www.zhenai.com//zhenghun//alashanmeng"}
	for index, city := range expectedCity {
		if city != expectedCity[index] {
			t.Errorf("Expected city #%d,%s;but had city %s",
				index, city, expectedCity[index])
		}
	}
	for index, url := range expectedUrl {
		if url != expectedUrl[index] {
			t.Errorf("Expected url #%d,%s;but had url %s",
				index, url, expectedCity[index])
		}
	}
}
