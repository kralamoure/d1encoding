// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type TutorialEnd struct{}

func (m TutorialEnd) ProtocolId() d1proto.MsgCliId {
	return d1proto.TutorialEnd
}

func (m TutorialEnd) Serialized() (string, error) {
	return "", nil
}

func (m *TutorialEnd) Deserialize(extra string) error {
	return nil
}
