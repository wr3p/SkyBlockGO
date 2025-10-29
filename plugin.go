package skyblockgo

import (
	"github.com/wr3p/SkyBlockGO/island"
)

type Plugin struct {
	IslandManager *island.Manager
}

func NewPlugin() *Plugin {
	p := &Plugin{}
	p.IslandManager = island.NewManager(p)
	return p
}
