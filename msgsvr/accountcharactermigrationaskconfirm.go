// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharacterMigrationAskConfirm struct{}

func (m AccountCharacterMigrationAskConfirm) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterMigrationAskConfirm
}

func (m AccountCharacterMigrationAskConfirm) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterMigrationAskConfirm) Deserialize(extra string) error {
	return nil
}
