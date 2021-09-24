package parser

type ParserResult struct {
	Item []interface{}
	Requests []Request
}

type Request struct{
	url string
	ParserFunc func([]byte) ParserResult
}