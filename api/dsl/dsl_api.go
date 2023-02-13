package dsl

type Request struct {
}

type Result struct {
}

type DSLMerger interface {
	Merge(*Request) *Result
}
