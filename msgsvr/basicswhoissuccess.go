// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsWhoIsSuccess struct{}

func (m BasicsWhoIsSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsWhoIsSuccess
}

func (m BasicsWhoIsSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsWhoIsSuccess) Deserialize(extra string) error {
	return nil
}
