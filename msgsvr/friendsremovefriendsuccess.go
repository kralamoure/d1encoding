// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FriendsRemoveFriendSuccess struct{}

func (m FriendsRemoveFriendSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FriendsRemoveFriendSuccess
}

func (m FriendsRemoveFriendSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *FriendsRemoveFriendSuccess) Deserialize(extra string) error {
	return nil
}
