// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeStopRepeatCraft struct{}

func (m ExchangeStopRepeatCraft) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeStopRepeatCraft
}

func (m ExchangeStopRepeatCraft) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeStopRepeatCraft) Deserialize(extra string) error {
	return nil
}
