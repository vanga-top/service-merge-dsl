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
