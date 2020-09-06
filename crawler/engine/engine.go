package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

//循环请求
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:] //拿出req的第一个元素
		log.Printf("Fectch %s", r.Url)
		//进行处理
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			//忽略它错
			log.Printf("拉取出错! url  %s : %v", r.Url, err)
			continue
		}
		//开始送parser
		parserResult := r.ParserFunc(body)
		requests = append(requests,
			parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}
