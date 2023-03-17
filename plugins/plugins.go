package plugins

type Status int

const (
	CREATED Status = iota
	INITED
	RUNNING
	STOP
	DESTROY
)

// Plugin for instance's plugins
type Plugin interface {
	ID() string
	Name() string
	Init() error
	Status() Status
	Destroy() error
}

/**
 	process schedule

	1、http request comes in
    2、gateway to parse http request --》 uri  to  service
    3、call remote call to get remote result
	4、call dsl to merge
*/