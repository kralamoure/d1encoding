// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMovementItems struct{}

func (m ExchangeMovementItems) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeMovementItems
}

func (m ExchangeMovementItems) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeMovementItems) Deserialize(extra string) error {
	return nil
}
