// automatically generated by stateify.

package timerfd

import (
	"github.com/dismantl/gvisor/pkg/state"
)

func (tfd *TimerFileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/timerfd.TimerFileDescription"
}

func (tfd *TimerFileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"events",
		"timer",
		"val",
	}
}

func (tfd *TimerFileDescription) beforeSave() {}

// +checklocksignore
func (tfd *TimerFileDescription) StateSave(stateSinkObject state.Sink) {
	tfd.beforeSave()
	stateSinkObject.Save(0, &tfd.vfsfd)
	stateSinkObject.Save(1, &tfd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &tfd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &tfd.NoLockFD)
	stateSinkObject.Save(4, &tfd.events)
	stateSinkObject.Save(5, &tfd.timer)
	stateSinkObject.Save(6, &tfd.val)
}

func (tfd *TimerFileDescription) afterLoad() {}

// +checklocksignore
func (tfd *TimerFileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &tfd.vfsfd)
	stateSourceObject.Load(1, &tfd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &tfd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &tfd.NoLockFD)
	stateSourceObject.Load(4, &tfd.events)
	stateSourceObject.Load(5, &tfd.timer)
	stateSourceObject.Load(6, &tfd.val)
}

func init() {
	state.Register((*TimerFileDescription)(nil))
}
