// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameGetMapData struct{}

func (m GameGetMapData) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameGetMapData
}

func (m GameGetMapData) Serialized() (string, error) {
	return "", nil
}

func (m *GameGetMapData) Deserialize(extra string) error {
	return nil
}
