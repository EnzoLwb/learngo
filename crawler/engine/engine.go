package engine

import (
	"learngo/crawler/fetcher"
	"learngo/crawler/persist"
	"log"
)

//循环请求
func Run(seeds ...Request) {
	var requests []Request
	var saveEngine saveEngineChan
	saveEngine.ItemChan = persist.ItemSaver()
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:] //拿出req的第一个元素
		//有单任务变成并发的worker  （fetch 和 parser 部分）
		parserResult, err := worker(r)
		if err != nil {
			continue //单个网络请求失败 也不要影响整个过程啊
		}
		requests = append(requests,
			parserResult.Requests...)
		for _, item := range parserResult.Items {
			go func() {
				saveEngine.ItemChan <- item //存储
			}()
		}
	}
}

//并发型的一个worker
func worker(r Request) (ParseResult, error) {
	log.Printf("Fectch %s", r.Url)
	//进行处理
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		//忽略它错
		log.Printf("拉取出错! url  %s : %v", r.Url, err)
		return ParseResult{}, nil
	}
	//开始送parser
	return r.ParserFunc(body), nil
}
