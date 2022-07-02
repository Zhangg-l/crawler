package parsers

import (
	"go_code/project11/crawler/model"
	"io/ioutil"
	"log"
	"testing"
)

func TestProfile(t *testing.T) {
	content, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		t.Errorf("ReadFile failed %v", err)
	}

	result := ParseFile(content, "ass")
	for _, v := range result.Item {
		if v, ok := v.(*model.Profile); ok {
			log.Printf("%#v", v)
		} else {
			t.Error("v not is (*model.Profile) type")
		}
	}

}
