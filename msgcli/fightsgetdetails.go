// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FightsGetDetails struct{}

func (m FightsGetDetails) ProtocolId() d1proto.MsgCliId {
	return d1proto.FightsGetDetails
}

func (m FightsGetDetails) Serialized() (string, error) {
	return "", nil
}

func (m *FightsGetDetails) Deserialize(extra string) error {
	return nil
}
