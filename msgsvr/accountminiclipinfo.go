// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountMiniClipInfo struct{}

func (m AccountMiniClipInfo) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountMiniClipInfo
}

func (m AccountMiniClipInfo) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountMiniClipInfo) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
