// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosMembers struct{}

func (m GuildInfosMembers) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildInfosMembers
}

func (m GuildInfosMembers) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosMembers) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
