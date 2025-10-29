package island

import "fmt"

type PluginRef interface{}

type Manager struct {
	plugin PluginRef
}

func NewManager(p PluginRef) *Manager {
	return &Manager{plugin: p}
}

func (m *Manager) Create(owner string) {
	fmt.Println("Creating island for", owner)
}
