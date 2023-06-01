// automatically generated by stateify.

package tcp

import (
	"github.com/dismantl/gvisor/pkg/state"
)

func (offset *TSOffset) StateTypeName() string {
	return "pkg/tcpip/internal/tcp.TSOffset"
}

func (offset *TSOffset) StateFields() []string {
	return []string{
		"milliseconds",
	}
}

func (offset *TSOffset) beforeSave() {}

// +checklocksignore
func (offset *TSOffset) StateSave(stateSinkObject state.Sink) {
	offset.beforeSave()
	stateSinkObject.Save(0, &offset.milliseconds)
}

func (offset *TSOffset) afterLoad() {}

// +checklocksignore
func (offset *TSOffset) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &offset.milliseconds)
}

func init() {
	state.Register((*TSOffset)(nil))
}
