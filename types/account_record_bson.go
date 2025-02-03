package types

import (
	bsonenc "github.com/ProtoconNet/mitum-currency/v3/digest/util/bson"
	ctypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (t AccountRecord) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(bson.M{
		"_hint":     t.Hint().String(),
		"address":   t.address,
		"amounts":   t.amounts,
		"last_time": t.lastTime,
	})
}

type AccountRecordBSONUnmarshaler struct {
	Hint     string            `bson:"_hint"`
	Address  string            `bson:"address"`
	Amounts  bson.Raw          `bson:"amounts"`
	LastTime map[string]uint64 `bson:"last_time"`
}

func (t *AccountRecord) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
	e := util.StringError("decode bson of AccountRecord")

	var u AccountRecordBSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	ht, err := hint.ParseHint(u.Hint)
	if err != nil {
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
	err = t.unpack(enc, ht, u.Address)
	if err != nil {
		return e.Wrap(err)
	}

	return nil
}
