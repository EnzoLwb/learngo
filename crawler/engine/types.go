package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

//这其实就是一个 解析器的结果
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//定义一个空的ParseResult 方便测试
func NilParseResult([]byte) ParseResult {
	return ParseResult{}
}
