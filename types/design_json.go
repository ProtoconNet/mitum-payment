package types

import (
	"encoding/json"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/pkg/errors"
)

type DesignJSONMarshaler struct {
	hint.BaseHinter
	Accounts map[string]AccountInfo `json:"accounts"`
}

func (de Design) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(DesignJSONMarshaler{
		BaseHinter: de.BaseHinter,
		Accounts:   de.accounts,
	})
}

type DesignJSONUnmarshaler struct {
	Hint     hint.Hint       `json:"_hint"`
	Accounts json.RawMessage `json:"accounts"`
}

func (de *Design) DecodeJSON(b []byte, enc encoder.Encoder) error {
	e := util.StringError("failed to decode json of Design")

	var u DesignJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	accounts := make(map[string]AccountInfo)
	m, err := enc.DecodeMap(u.Accounts)
	if err != nil {
		return e.Wrap(err)
	}
	for k, v := range m {
		ac, ok := v.(AccountInfo)
		if !ok {
			return e.Wrap(errors.Errorf("expected AccountInfo, not %T", v))
		}

		accounts[k] = ac
	}
	de.accounts = accounts

	err = de.unpack(enc, u.Hint)
	if err != nil {
		return e.Wrap(err)
	}

	return nil
}
