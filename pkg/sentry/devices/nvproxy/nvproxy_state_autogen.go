// automatically generated by stateify.

package nvproxy

import (
	"github.com/dismantl/gvisor/pkg/state"
)

func (dev *frontendDevice) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.frontendDevice"
}

func (dev *frontendDevice) StateFields() []string {
	return []string{
		"nvp",
		"minor",
	}
}

func (dev *frontendDevice) beforeSave() {}

// +checklocksignore
func (dev *frontendDevice) StateSave(stateSinkObject state.Sink) {
	dev.beforeSave()
	stateSinkObject.Save(0, &dev.nvp)
	stateSinkObject.Save(1, &dev.minor)
}

func (dev *frontendDevice) afterLoad() {}

// +checklocksignore
func (dev *frontendDevice) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &dev.nvp)
	stateSourceObject.Load(1, &dev.minor)
}

func (n *nvproxy) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.nvproxy"
}

func (n *nvproxy) StateFields() []string {
	return []string{
		"version",
	}
}

// +checklocksignore
func (n *nvproxy) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.version)
}

// +checklocksignore
func (n *nvproxy) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.version)
	stateSourceObject.AfterLoad(n.afterLoad)
}

func (o *object) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.object"
}

func (o *object) StateFields() []string {
	return []string{
		"impl",
	}
}

func (o *object) beforeSave() {}

// +checklocksignore
func (o *object) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.impl)
}

func (o *object) afterLoad() {}

// +checklocksignore
func (o *object) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.impl)
}

func (o *osDescMem) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.osDescMem"
}

func (o *osDescMem) StateFields() []string {
	return []string{
		"object",
		"pinnedRanges",
	}
}

func (o *osDescMem) beforeSave() {}

// +checklocksignore
func (o *osDescMem) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.object)
	stateSinkObject.Save(1, &o.pinnedRanges)
}

func (o *osDescMem) afterLoad() {}

// +checklocksignore
func (o *osDescMem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.object)
	stateSourceObject.Load(1, &o.pinnedRanges)
}

func (dev *uvmDevice) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.uvmDevice"
}

func (dev *uvmDevice) StateFields() []string {
	return []string{
		"nvp",
	}
}

func (dev *uvmDevice) beforeSave() {}

// +checklocksignore
func (dev *uvmDevice) StateSave(stateSinkObject state.Sink) {
	dev.beforeSave()
	stateSinkObject.Save(0, &dev.nvp)
}

func (dev *uvmDevice) afterLoad() {}

// +checklocksignore
func (dev *uvmDevice) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &dev.nvp)
}

func (v *DriverVersion) StateTypeName() string {
	return "pkg/sentry/devices/nvproxy.DriverVersion"
}

func (v *DriverVersion) StateFields() []string {
	return []string{
		"major",
		"minor",
		"patch",
	}
}

func (v *DriverVersion) beforeSave() {}

// +checklocksignore
func (v *DriverVersion) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	stateSinkObject.Save(0, &v.major)
	stateSinkObject.Save(1, &v.minor)
	stateSinkObject.Save(2, &v.patch)
}

func (v *DriverVersion) afterLoad() {}

// +checklocksignore
func (v *DriverVersion) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.major)
	stateSourceObject.Load(1, &v.minor)
	stateSourceObject.Load(2, &v.patch)
}

func init() {
	state.Register((*frontendDevice)(nil))
	state.Register((*nvproxy)(nil))
	state.Register((*object)(nil))
	state.Register((*osDescMem)(nil))
	state.Register((*uvmDevice)(nil))
	state.Register((*DriverVersion)(nil))
}
