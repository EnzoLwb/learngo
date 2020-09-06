package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" [^>]+>([^<]+)</a>`

//输入原网页 返回解析后的子字符串
func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(CityListRe)

	//建立解析结构体
	result := engine.ParseResult{}              //数组结构体
	matches := re.FindAllSubmatch(contents, -1) //返回 [][]string
	limit := 10                                 //限制爬取city页面的次数 因为要爬取人之前是要等所有city页面结束的
	for _, i := range matches {
		//返回
		result.Items = append(result.Items, string(i[2])) //放入city名称
		newRequests := engine.Request{
			Url:        string(i[1]),
			ParserFunc: ParserCity, //这是解析每个城市页面里面人的解析器 也就是迷宫算法里的一步步走往深了走
			//engine.NilParseResult 不能加括号 因为添加括号表示返回函数体的返回值 ，而这里定义的是返回一个函数
		}
		result.Requests = append(result.Requests, newRequests)
		//fmt.Printf("City:%s Url:%s", i[2], i[1])
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
