// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedCommandPrompt struct{}

func (m BasicsAuthorizedCommandPrompt) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsAuthorizedCommandPrompt
}

func (m BasicsAuthorizedCommandPrompt) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedCommandPrompt) Deserialize(extra string) error {
	return nil
}
