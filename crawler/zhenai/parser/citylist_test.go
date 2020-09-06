package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("crawler/zhenai/parser/citylist.html")
	if err != nil {
		panic(err)
	}
	results := ParserCityList(contents)
	fmt.Println(len(results.Requests))
}
