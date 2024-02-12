// automatically generated by stateify.

package hostinet

import (
	"github.com/dismantl/gvisor/pkg/state"
)

func (s *Socket) StateTypeName() string {
	return "pkg/sentry/socket/hostinet.Socket"
}

func (s *Socket) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
		"DentryMetadataFileDescriptionImpl",
		"SendReceiveTimeout",
		"family",
		"stype",
		"protocol",
		"queue",
		"fd",
		"recvClosed",
	}
}

func (s *Socket) beforeSave() {}

// +checklocksignore
func (s *Socket) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.vfsfd)
	stateSinkObject.Save(1, &s.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &s.LockFD)
	stateSinkObject.Save(3, &s.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(4, &s.SendReceiveTimeout)
	stateSinkObject.Save(5, &s.family)
	stateSinkObject.Save(6, &s.stype)
	stateSinkObject.Save(7, &s.protocol)
	stateSinkObject.Save(8, &s.queue)
	stateSinkObject.Save(9, &s.fd)
	stateSinkObject.Save(10, &s.recvClosed)
}

func (s *Socket) afterLoad() {}

// +checklocksignore
func (s *Socket) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.vfsfd)
	stateSourceObject.Load(1, &s.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &s.LockFD)
	stateSourceObject.Load(3, &s.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(4, &s.SendReceiveTimeout)
	stateSourceObject.Load(5, &s.family)
	stateSourceObject.Load(6, &s.stype)
	stateSourceObject.Load(7, &s.protocol)
	stateSourceObject.Load(8, &s.queue)
	stateSourceObject.Load(9, &s.fd)
	stateSourceObject.Load(10, &s.recvClosed)
}

func init() {
	state.Register((*Socket)(nil))
}
