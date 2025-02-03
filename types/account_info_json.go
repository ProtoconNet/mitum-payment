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

type AccountInfoJSONMarshaler struct {
	hint.BaseHinter
	Address       base.Address             `json:"address"`
	TransferLimit map[string]ctypes.Amount `json:"transfer_limit"`
	PeriodTime    map[string][3]uint64     `json:"period_time"`
}

func (t AccountInfo) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(AccountInfoJSONMarshaler{
		BaseHinter:    t.BaseHinter,
		Address:       t.address,
		TransferLimit: t.transferLimit,
		PeriodTime:    t.periodTime,
	})
}

type AccountInfoJSONUnmarshaler struct {
	Hint          hint.Hint            `json:"_hint"`
	Address       string               `json:"address"`
	TransferLimit json.RawMessage      `json:"transfer_limit"`
	PeriodTime    map[string][3]uint64 `json:"period_time"`
}

func (t *AccountInfo) DecodeJSON(b []byte, enc encoder.Encoder) error {
	e := util.StringError("failed to decode json of AccountInfo")

	var u AccountInfoJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	tLimit := make(map[string]ctypes.Amount)
	m, err := enc.DecodeMap(u.TransferLimit)
	if err != nil {
		return e.Wrap(err)
	}
	for k, v := range m {
		am, ok := v.(ctypes.Amount)
		if !ok {
			return e.Wrap(errors.Errorf("expected Amount, not %T", v))
		}

		tLimit[k] = am
	}
	t.transferLimit = tLimit
	t.periodTime = u.PeriodTime

	err = t.unpack(enc, u.Hint, u.Address)
	if err != nil {
		return e.Wrap(err)
	}

	return nil
}
