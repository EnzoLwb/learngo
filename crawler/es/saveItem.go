package main

import (
	"context"
	"encoding/json"
	"fmt"
	"learngo/crawler/model"

	"github.com/olivere/elastic/v7"
)

//存数据
func create(item model.Profile) (string, error) {
	//set url 可以设置服务器地址 默认就是本机9200端口
	cli, err := elastic.NewClient(elastic.SetSniff(false)) //in docker must turn off
	if err != nil {
		return "", err
	}
	res, err := cli.Index().Index(Index).BodyJson(item).Do(context.Background()) //三元组 index 代表数据库名 type 已经被移除可以理解为table，id 这里不写 他可以自动生成
	if err != nil {
		return "", err
	}
	return res.Id, nil
}

//拿数据
func getById(id string) {
	cli, err := elastic.NewClient(elastic.SetSniff(false)) //in docker must turn off
	if err != nil {
		panic(err)
	}
	res, err := cli.Get().Index(Index).Id(id).Do(context.Background())
	//fmt.Printf("%+v", res)
	//转序列化
	var actual model.Profile
	err = json.Unmarshal(res.Source, &actual)
	if err != nil {
		panic(err)
	}
	fmt.Println("actual:")
	fmt.Printf("%+v", actual)
}

const Index = "ikes2"

func main2() {
	/*item := model.Profile{
		Name:    "enzo",
		Hunyin:  "离婚",
		Age:     "25",
		Xingzuo: "双子座",
	}*/
	//id, _ := create(item)
	//getById(id)
	//查数据
}
