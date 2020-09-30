package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
)

//创建 index curl -XPUT http://localhost:9200/index
const url = "http://localhost:9200/"

func createIndex(indexName string) {
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodPut, url+indexName, nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//不要头部的信息 只需要判code信息即可
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("error:status code %d", resp.StatusCode) //fmt.Errorf()
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(all)) //{"acknowledged":true,"shards_acknowledged":true,"index":"ikes2"}
}

/* 创建映射  配置字段映射
curl -XPOST http://localhost:9200/index/_mapping -H 'Content-Type:application/json' -d'
{
        "properties": {
            "content": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            }
        }
}
*/
func mapping(indexName string) {
	// json
	contentType := "application/json"
	data := `{
        "properties": {
            "content": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
			"ep": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            }
        }
	}`
	resp, err := http.Post(url+indexName+"/_mapping", contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b)) //{"acknowledged":true}
}

/* 录入数据
curl -XPOST http://localhost:9200/index/_create/1 -H 'Content-Type:application/json' -d'
{"content":"美国留给伊拉克的是个烂摊子吗"}
*/
func inputData(indexName string, data string, id string) {
	// json
	contentType := "application/json"
	var inputUrl string
	if id != "0" {
		inputUrl = url + indexName + "/_create/" + id
	} else {
		inputUrl = url + indexName + "/_doc"
	}
	resp, err := http.Post(inputUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b)) //{"acknowledged":true}
	//{"_index":"ikes2","_type":"_doc","_id":"JKZWbHQB8wLP38GV5ldD","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":2,"_primary_term":1}
}

//测试分词效果
//curl -X POST "http://localhost:9200/your_index/_analyze?pretty" -H 'Content-Type: application/json' -d'
func analyze(indexName string) {
	// json
	contentType := "application/json"
	data := `{
		  "analyzer": "ik_max_word",
		  "text":     "laravel天下无敌"
		}`
	resp, err := http.Post(url+indexName+"/_analyze", contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b)) //{"tokens":[{"token":"laravel","start_offset":0,"end_offset":7,"type":"ENGLISH","position":0},{"token":"天下无敌","start_offset":7,"end_offset":11,"type":"CN_WORD","position":1},{"token":"天下","start_offset":7,"end_offset":9,"type":"CN_WORD","position":2},{"token":"无敌","start_offset":9,"end_offset":11,"type":"CN_WORD","position":3}]}
}

/* 搜索
curl -XPOST http://localhost:9200/index/_search  -H 'Content-Type:application/json' -d'
{
    "query" : { "match" : { "content" : "中国" }},
    "highlight" : {
        "pre_tags" : ["<tag1>", "<tag2>"],
        "post_tags" : ["</tag1>", "</tag2>"],
        "fields" : {
            "content" : {}
        }
    }
}*/
func highLightSearch(indexName string, search string) {
	// json
	contentType := "application/json"
	data := `{
		"query" : { "match" : { "content" : "` + search + `" }},
		"highlight" : {
			"pre_tags" : ["<tag1>", "<tag2>"],
			"post_tags" : ["</tag1>", "</tag2>"],
			"fields" : {
				"content" : {}
			}
		}
	}`
	resp, err := http.Post(url+indexName+"/_search", contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

//http://localhost:9200/ikes2/_search?q=ep:3333
func search(indexName string, attr string, search string) {
	var attrSearchUrl string
	if attr != "" {
		attrSearchUrl = attr + ":" + search
	} else {
		attrSearchUrl = search
	}
	resp, err := http.Get(url + indexName + "/_search?q=" + url2.QueryEscape(attrSearchUrl))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}

	fmt.Println(string(b))
}
func main() {
	createIndex("ikes2")
	mapping("ikes2")
	analyze("ikes2")
	inputData("ikes2", `{"content":"明天星期几","ep":"今天是星期三"}`, "0")
	highLightSearch("ikes2", "烂渔船")
	//search("ikes2", "ep", "今天")
}
