// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismInfosClosing struct{}

func (m ConquestPrismInfosClosing) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismInfosClosing
}

func (m ConquestPrismInfosClosing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismInfosClosing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
