// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type SubwayPrismUse struct{}

func (m SubwayPrismUse) ProtocolId() retroproto.MsgCliId {
	return retroproto.SubwayPrismUse
}

func (m SubwayPrismUse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayPrismUse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
