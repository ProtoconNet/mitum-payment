package types

import (
	"encoding/json"
	ctypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/ProtoconNet/mitum2/util/valuehash"
)

var AccountRecordHint = hint.MustNewHint("mitum-payment-account-record-v0.0.1")

type AccountRecord struct {
	hint.BaseHinter
	address  base.Address
	amounts  map[string]ctypes.Amount
	lastTime map[string]uint64
}

func NewAccountRecord(
	address base.Address,
) AccountRecord {
	amounts := make(map[string]ctypes.Amount)
	lastTime := make(map[string]uint64)
	return AccountRecord{
		BaseHinter: hint.NewBaseHinter(AccountRecordHint),
		address:    address,
		amounts:    amounts,
		lastTime:   lastTime,
	}
}

func NewEmptyAccountRecord() AccountRecord {
	return AccountRecord{
		BaseHinter: hint.NewBaseHinter(AccountRecordHint),
	}
}

func (t AccountRecord) IsValid([]byte) error {
	if err := t.BaseHinter.IsValid(nil); err != nil {
		return err
	}

	if err := util.CheckIsValiders(nil, false,
		t.address,
	); err != nil {
		return err
	}

	for _, v := range t.amounts {
		if err := util.CheckIsValiders(nil, false,
			v,
		); err != nil {
			return err
		}
	}

	return nil
}

func (t AccountRecord) Bytes() []byte {
	var bam []byte
	if t.amounts != nil {
		am, _ := json.Marshal(t.amounts)
		bam = valuehash.NewSHA256(am).Bytes()
	} else {
		bam = []byte{}
	}

	var blt []byte
	if t.lastTime != nil {
		lt, _ := json.Marshal(t.lastTime)
		blt = valuehash.NewSHA256(lt).Bytes()
	} else {
		blt = []byte{}
	}

	return util.ConcatBytesSlice(
		t.address.Bytes(),
		bam,
		blt,
	)
}

func (t AccountRecord) Address() base.Address {
	return t.address
}

func (t *AccountRecord) SetAmount(cid string, am ctypes.Amount) {
	t.amounts[cid] = am
}

func (t AccountRecord) Amount(cid string) *ctypes.Amount {
	am, found := t.amounts[cid]
	if !found {
		return nil
	}

	return &am
}

func (t *AccountRecord) SetLastTime(cid string, lastTime uint64) {
	t.lastTime[cid] = lastTime
}

func (t AccountRecord) LastTime(cid string) *uint64 {
	lt, found := t.lastTime[cid]
	if !found {
		return nil
	}

	return &lt
}
