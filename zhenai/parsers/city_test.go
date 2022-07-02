package parsers

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestCity(t *testing.T) {
	// 获取数据
	contents, err := ioutil.ReadFile("city_test_data.html")

	if err != nil {
		panic(fmt.Sprintf("file parse fail %s", err))
	}
	result := ParseCity(contents)

	for i, v := range result.Item {
		if v, ok := v.(string); ok {
			log.Printf("%v : URL:%s\n", v, result.Requests[i].Url)
		} else {
			t.Error("v not is (*model.Profile) type")
		}
	}
}
