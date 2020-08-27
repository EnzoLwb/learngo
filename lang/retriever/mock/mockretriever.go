package mock

import "fmt"

//假设是项目中的测试代码
type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf(
		"mockRetriever: {Contents=%s}", r.Contents)
}

func (r *Retriever) Post(url string,
	form map[string]string) string {
	r.Contents = form["Contents"]
	return "ok"
}
func (r *Retriever) Get(url string) string {
	return r.Contents
}
