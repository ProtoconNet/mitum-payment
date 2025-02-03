package state

import (
	bsonenc "github.com/ProtoconNet/mitum-currency/v3/digest/util/bson"
	"github.com/ProtoconNet/mitum-payment/types"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"go.mongodb.org/mongo-driver/bson"
)

func (sv DesignStateValue) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(
		bson.M{
			"_hint":  sv.Hint().String(),
			"design": sv.Design,
		},
	)
}

type DesignStateValueBSONUnmarshaler struct {
	Hint   string   `bson:"_hint"`
	Design bson.Raw `bson:"design"`
}

func (sv *DesignStateValue) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
	e := util.StringError("decode bson of DesignStateValue")

	var u DesignStateValueBSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	ht, err := hint.ParseHint(u.Hint)
	if err != nil {
		return e.Wrap(err)
	}
	sv.BaseHinter = hint.NewBaseHinter(ht)

	var sd types.Design
	if err := sd.DecodeBSON(u.Design, enc); err != nil {
		return e.Wrap(err)
	}
	sv.Design = sd

	return nil
}

func (sv AccountRecordStateValue) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(
		bson.M{
			"_hint":          sv.Hint().String(),
			"account_record": sv.AccountRecord,
		},
	)
}

type AccountRecordStateValueBSONUnmarshaler struct {
	Hint          string   `bson:"_hint"`
	AccountRecord bson.Raw `bson:"account_record"`
}

func (sv *AccountRecordStateValue) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
	e := util.StringError("decode bson of AccountRecordStateValue")

	var u AccountRecordStateValueBSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	ht, err := hint.ParseHint(u.Hint)
	if err != nil {
		return e.Wrap(err)
	}
	sv.BaseHinter = hint.NewBaseHinter(ht)

	var accountRecord types.AccountRecord
	if err := accountRecord.DecodeBSON(u.AccountRecord, enc); err != nil {
		return e.Wrap(err)
	}
	sv.AccountRecord = accountRecord

	return nil
}

//func (sv ItemStateValue) MarshalBSON() ([]byte, error) {
//	return bsonenc.Marshal(
//		bson.M{
//			"_hint":          sv.Hint().String(),
//			"timestamp_item": sv.Item,
//		},
//	)
//}
//
//type ItemStateValueBSONUnmarshaler struct {
//	Hint          string   `bson:"_hint"`
//	TimeStampItem bson.Raw `bson:"timestamp_item"`
//}
//
//func (sv *ItemStateValue) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
//	e := util.StringError("decode bson of ItemStateValue")
//
//	var u ItemStateValueBSONUnmarshaler
//	if err := enc.Unmarshal(b, &u); err != nil {
//		return e.Wrap(err)
//	}
//
//	ht, err := hint.ParseHint(u.Hint)
//	if err != nil {
//		return e.Wrap(err)
//	}
//	sv.BaseHinter = hint.NewBaseHinter(ht)
//
//	var n types.AccountInfo
//	if err := n.DecodeBSON(u.TimeStampItem, enc); err != nil {
//		return e.Wrap(err)
//	}
//	sv.Item = n
//
//	return nil
//}
