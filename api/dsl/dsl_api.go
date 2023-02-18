package dsl

type Request struct {
}

type Result struct {
}

type DSLMerger interface {
	Merge(*Request) *Result
	Parse(filePath string) (*Request, *Result) //根据文件直接解析出请求内容
}
