// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountEquipAdd struct{}

func (m MountEquipAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountEquipAdd
}

func (m MountEquipAdd) Serialized() (string, error) {
	return "", nil
}

func (m *MountEquipAdd) Deserialize(extra string) error {
	return nil
}
