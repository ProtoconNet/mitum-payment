package state

import (
	"encoding/json"

	"github.com/ProtoconNet/mitum-payment/types"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/ProtoconNet/mitum2/util/hint"
)

type DesignStateValueJSONMarshaler struct {
	hint.BaseHinter
	Design types.Design `json:"design"`
}

func (sv DesignStateValue) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(
		DesignStateValueJSONMarshaler(sv),
	)
}

type DesignStateValueJSONUnmarshaler struct {
	Hint   hint.Hint       `json:"_hint"`
	Design json.RawMessage `json:"design"`
}

func (sv *DesignStateValue) DecodeJSON(b []byte, enc encoder.Encoder) error {
	e := util.StringError("failed to decode json of DesignStateValue")

	var u DesignStateValueJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	sv.BaseHinter = hint.NewBaseHinter(u.Hint)

	var sd types.Design
	if err := sd.DecodeJSON(u.Design, enc); err != nil {
		return e.Wrap(err)
	}
	sv.Design = sd

	return nil
}

type AccountRecordStateValueJSONMarshaler struct {
	hint.BaseHinter
	AccountRecord types.AccountRecord `json:"account_record"`
}

func (sv AccountRecordStateValue) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(
		AccountRecordStateValueJSONMarshaler(sv),
	)
}

type AccountRecordStateValueJSONUnmarshaler struct {
	Hint          hint.Hint       `json:"_hint"`
	AccountRecord json.RawMessage `json:"account_record"`
}

func (sv *AccountRecordStateValue) DecodeJSON(b []byte, enc encoder.Encoder) error {
	e := util.StringError("failed to decode json of AccountRecordStateValue")

	var u AccountRecordStateValueJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	sv.BaseHinter = hint.NewBaseHinter(u.Hint)
	var accountRecord types.AccountRecord
	if err := accountRecord.DecodeJSON(u.AccountRecord, enc); err != nil {
		return e.Wrap(err)
	}
	sv.AccountRecord = accountRecord

	return nil
}

//type TimeStampItemStateValueJSONMarshaler struct {
//	hint.BaseHinter
//	Item types.AccountInfo `json:"timestamp_item"`
//}
//
//func (sv ItemStateValue) MarshalJSON() ([]byte, error) {
//	return util.MarshalJSON(
//		TimeStampItemStateValueJSONMarshaler(sv),
//	)
//}
//
//type ItemStateValueJSONUnmarshaler struct {
//	Hint          hint.Hint       `json:"_hint"`
//	TimeStampItem json.RawMessage `json:"timestamp_item"`
//}
//
//func (sv *ItemStateValue) DecodeJSON(b []byte, enc encoder.Encoder) error {
//	e := util.StringError("decode json of ItemStateValue")
//
//	var u ItemStateValueJSONUnmarshaler
//	if err := enc.Unmarshal(b, &u); err != nil {
//		return e.Wrap(err)
//	}
//
//	sv.BaseHinter = hint.NewBaseHinter(u.Hint)
//
//	var t types.AccountInfo
//	if err := t.DecodeJSON(u.TimeStampItem, enc); err != nil {
//		return e.Wrap(err)
//	}
//	sv.Item = t
//
//	return nil
//}
