// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeAccept struct{}

func (m ExchangeAccept) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeAccept
}

func (m ExchangeAccept) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeAccept) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
