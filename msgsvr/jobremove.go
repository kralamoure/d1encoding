// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type JobRemove struct{}

func (m JobRemove) ProtocolId() retroproto.MsgSvrId {
	return retroproto.JobRemove
}

func (m JobRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
