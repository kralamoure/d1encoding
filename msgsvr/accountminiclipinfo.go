// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountMiniClipInfo struct{}

func (m AccountMiniClipInfo) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountMiniClipInfo
}

func (m AccountMiniClipInfo) Serialized() (string, error) {
	return "", nil
}

func (m *AccountMiniClipInfo) Deserialize(extra string) error {
	return nil
}
