package types

import (
	bsonenc "github.com/ProtoconNet/mitum-currency/v3/digest/util/bson"
	ctypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (t AccountInfo) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(bson.M{
		"_hint":          t.Hint().String(),
		"address":        t.address,
		"transfer_limit": t.transferLimit,
		"period_time":    t.periodTime,
	})
}

type AccountInfoBSONUnmarshaler struct {
	Hint          string               `bson:"_hint"`
	Address       string               `bson:"address"`
	TransferLimit bson.Raw             `bson:"transfer_limit"`
	PeriodTime    map[string][3]uint64 `bson:"period_time"`
}

func (t *AccountInfo) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
	e := util.StringError("decode bson of AccountInfo")

	var u AccountInfoBSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	ht, err := hint.ParseHint(u.Hint)
	if err != nil {
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

	err = t.unpack(enc, ht, u.Address)
	if err != nil {
		return e.Wrap(err)
	}

	return nil
}
