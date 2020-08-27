package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

//假设这是一个爬虫类
type Retriever struct {
	UserAgent string //public 首字母大写 private 首字母小写
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(
		resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
