// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildInfosTaxCollectorsMovement struct{}

func (m GuildInfosTaxCollectorsMovement) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildInfosTaxCollectorsMovement
}

func (m GuildInfosTaxCollectorsMovement) Serialized() (string, error) {
	return "", nil
}

func (m *GuildInfosTaxCollectorsMovement) Deserialize(extra string) error {
	return nil
}
