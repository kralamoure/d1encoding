// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GildUserInterfaceOpen struct{}

func (m GildUserInterfaceOpen) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GildUserInterfaceOpen
}

func (m GildUserInterfaceOpen) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildUserInterfaceOpen) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
