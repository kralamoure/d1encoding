// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesListRemove struct{}

func (m HousesListRemove) ProtocolId() retroproto.MsgSvrId {
	return retroproto.HousesListRemove
}

func (m HousesListRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesListRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
