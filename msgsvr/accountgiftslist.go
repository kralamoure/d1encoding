// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountGiftsList struct{}

func (m AccountGiftsList) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountGiftsList
}

func (m AccountGiftsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountGiftsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
