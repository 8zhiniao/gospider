package engine

type PaseResult struct {
	Requests []Request       //解析器返回一个url对应的解析器名称
	Items []interface{}     //解析器对应的结果值
}

type Request struct{
	Url string               //下一次要爬虫的地址
	ParserFunc func([]byte) PaseResult //地址内容对应的解析器
}

func NilParser([]byte) PaseResult {
	return PaseResult{}
}