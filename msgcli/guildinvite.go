// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildInvite struct{}

func (m GuildInvite) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildInvite
}

func (m GuildInvite) Serialized() (string, error) {
	return "", nil
}

func (m *GuildInvite) Deserialize(extra string) error {
	return nil
}
