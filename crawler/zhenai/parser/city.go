package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

//<th><a href="http://album.zhenai.com/u/1807074256" target="_blank">余生有你</a></th>
const CityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

//输入原网页 返回解析后的子字符串
func ParserCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(CityRe)

	//建立解析结构体
	result := engine.ParseResult{}              //数组结构体
	matches := re.FindAllSubmatch(contents, -1) //返回 [][]string
	for _, i := range matches {
		//返回
		name := string(i[2])                              //需要重新定义一个name 不然就会是重复的city页面里面的名称。因为request是在for循环后执行
		result.Items = append(result.Items, "User:"+name) //放入city名称
		newRequests := engine.Request{
			Url: string(i[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParserProfile(c, name)
			},
		}
		result.Requests = append(result.Requests, newRequests)
		//fmt.Printf("City:%s Url:%s", i[2], i[1])
	}
	return result
}
