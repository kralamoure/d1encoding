package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type AccountLoginSuccess struct {
	Authorized bool
}

func (m AccountLoginSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountLoginSuccess
}

func (m AccountLoginSuccess) Serialized() (string, error) {
	var authorized int
	if m.Authorized {
		authorized = 1
	}

	return fmt.Sprintf("%d", authorized), nil
}

func (m *AccountLoginSuccess) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1proto.ErrInvalidMsg
	}

	authorized, err := strconv.ParseBool(extra)
	if err != nil {
		return err
	}
	m.Authorized = authorized

	return nil
}
