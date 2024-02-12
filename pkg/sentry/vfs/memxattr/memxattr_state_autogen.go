// automatically generated by stateify.

package memxattr

import (
	"github.com/dismantl/gvisor/pkg/state"
)

func (x *SimpleExtendedAttributes) StateTypeName() string {
	return "pkg/sentry/vfs/memxattr.SimpleExtendedAttributes"
}

func (x *SimpleExtendedAttributes) StateFields() []string {
	return []string{
		"xattrs",
	}
}

func (x *SimpleExtendedAttributes) beforeSave() {}

// +checklocksignore
func (x *SimpleExtendedAttributes) StateSave(stateSinkObject state.Sink) {
	x.beforeSave()
	stateSinkObject.Save(0, &x.xattrs)
}

func (x *SimpleExtendedAttributes) afterLoad() {}

// +checklocksignore
func (x *SimpleExtendedAttributes) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &x.xattrs)
}

func init() {
	state.Register((*SimpleExtendedAttributes)(nil))
}
