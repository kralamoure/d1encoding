// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeSellSuccess struct{}

func (m ExchangeSellSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeSellSuccess
}

func (m ExchangeSellSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeSellSuccess) Deserialize(extra string) error {
	return nil
}
