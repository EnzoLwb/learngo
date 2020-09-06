package parser

import (
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
)

const profileRe = `<div class="m-btn purple"[^>]*>([^<]*)</div>`

//输入原网页 返回解析后的子字符串
func ParserProfile(contents []byte,
	name string) engine.ParseResult {
	re := regexp.MustCompile(profileRe)
	matches := re.FindAllSubmatch(contents, -1) //返回 [][]string
	var profile model.Profile

	/*fmt.Println(string(matches[0][0])) //<div class="m-btn purple" data-v-8b1eac0c>离异</div>
	fmt.Println(string(matches[0][1])) //离异
	fmt.Println(string(matches[1][0])) //<div class="m-btn purple" data-v-8b1eac0c>63岁</div>
	*/
	if len(matches) > 0 {
		profile.Name = name
		profile.Hunyin = string(matches[0][1])
		profile.Age = string(matches[1][1])
		profile.Xingzuo = string(matches[2][1])
	} else {
		fmt.Println("未找到用户信息")
	}

	//建立解析结构体
	result := engine.ParseResult{
		Items: []interface{}{profile},
		//结束了 没有新的requests
	}
	fmt.Println(result.Items)
	return result
}
