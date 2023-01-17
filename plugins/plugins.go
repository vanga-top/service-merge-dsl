package plugins

// Plugin for instance's plugins
type Plugin interface {
	ID() string
	Init()
}
