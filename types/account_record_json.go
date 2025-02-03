package types

import (
	"encoding/json"
	ctypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/pkg/errors"
)

type AccountRecordJSONMarshaler struct {
	hint.BaseHinter
	Address  base.Address             `json:"address"`
	Amounts  map[string]ctypes.Amount `json:"amounts"`
	LastTime map[string]uint64        `json:"last_time"`
}

func (t AccountRecord) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(AccountRecordJSONMarshaler{
		BaseHinter: t.BaseHinter,
		Address:    t.address,
		Amounts:    t.amounts,
		LastTime:   t.lastTime,
	})
}

type AccountRecordJSONUnmarshaler struct {
	Hint     hint.Hint         `json:"_hint"`
	Address  string            `json:"address"`
	Amounts  json.RawMessage   `json:"amounts"`
	LastTime map[string]uint64 `json:"last_time"`
}

func (t *AccountRecord) DecodeJSON(b []byte, enc encoder.Encoder) error {
	e := util.StringError("failed to decode json of AccountRecord")

	var u AccountRecordJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	amounts := make(map[string]ctypes.Amount)
	m, err := enc.DecodeMap(u.Amounts)
	if err != nil {
		return e.Wrap(err)
	}
	for k, v := range m {
		am, ok := v.(ctypes.Amount)
		if !ok {
			return e.Wrap(errors.Errorf("expected Amount, not %T", v))
		}

		amounts[k] = am
	}
	t.amounts = amounts
	t.lastTime = u.LastTime
	err = t.unpack(enc, u.Hint, u.Address)
	if err != nil {
		return e.Wrap(err)
	}

	return nil
}
