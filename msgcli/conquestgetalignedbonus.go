// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestGetAlignedBonus struct{}

func (m ConquestGetAlignedBonus) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestGetAlignedBonus
}

func (m ConquestGetAlignedBonus) Serialized() (string, error) {
	return "", nil
}

func (m *ConquestGetAlignedBonus) Deserialize(extra string) error {
	return nil
}
