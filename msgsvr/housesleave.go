// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesLeave struct{}

func (m HousesLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesLeave
}

func (m HousesLeave) Serialized() (string, error) {
	return "", nil
}

func (m *HousesLeave) Deserialize(extra string) error {
	return nil
}
