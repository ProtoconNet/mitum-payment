package types

import (
	"encoding/json"
	"github.com/ProtoconNet/mitum-currency/v3/common"
	ctypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/ProtoconNet/mitum2/util/valuehash"
)

var AccountInfoHint = hint.MustNewHint("mitum-payment-account-info-v0.0.1")

type AccountInfo struct {
	hint.BaseHinter
	address       base.Address
	transferLimit map[string]ctypes.Amount
	periodTime    map[string][3]uint64 // startTime, endTime, duration
}

func NewAccountInfo(
	address base.Address,
) AccountInfo {
	tLimit := make(map[string]ctypes.Amount)
	pTime := make(map[string][3]uint64)
	return AccountInfo{
		BaseHinter:    hint.NewBaseHinter(AccountInfoHint),
		address:       address,
		transferLimit: tLimit,
		periodTime:    pTime,
	}
}

func (t AccountInfo) IsValid([]byte) error {
	if err := t.BaseHinter.IsValid(nil); err != nil {
		return err
	}

	if err := util.CheckIsValiders(nil, false,
		t.address,
	); err != nil {
		return err
	}

	for _, v := range t.transferLimit {
		if err := util.CheckIsValiders(nil, false,
			v,
		); err != nil {
			return err
		}
	}

	for _, v := range t.periodTime {
		if v[0] < 1 {
			return common.ErrFactInvalid.Wrap(common.ErrValueInvalid.Errorf("start time must be bigger than zero"))
		} else if v[1] < 1 {
			return common.ErrFactInvalid.Wrap(common.ErrValueInvalid.Errorf("end time must be bigger than zero"))
		} else if v[2] < 1 {
			return common.ErrFactInvalid.Wrap(common.ErrValueInvalid.Errorf("duration must be bigger than zero"))
		}
	}

	return nil
}

func (t AccountInfo) Bytes() []byte {
	var btl []byte
	if t.transferLimit != nil {
		tl, _ := json.Marshal(t.transferLimit)
		btl = valuehash.NewSHA256(tl).Bytes()
	} else {
		btl = []byte{}
	}

	var bpt []byte
	if t.periodTime != nil {
		pt, _ := json.Marshal(t.periodTime)
		bpt = valuehash.NewSHA256(pt).Bytes()
	} else {
		bpt = []byte{}
	}

	return util.ConcatBytesSlice(
		t.address.Bytes(),
		btl,
		bpt,
	)
}

func (t AccountInfo) Address() base.Address {
	return t.address
}

func (t *AccountInfo) SetTransferLimit(tLimit ctypes.Amount) {
	t.transferLimit[tLimit.Currency().String()] = tLimit
}

func (t AccountInfo) TransferLimit(cid string) *ctypes.Amount {
	am, found := t.transferLimit[cid]
	if !found {
		return nil
	}

	return &am
}

func (t AccountInfo) PeriodTime(cid string) *[3]uint64 {
	pt, found := t.periodTime[cid]
	if !found {
		return nil
	}

	return &pt
}

func (t *AccountInfo) SetPeriodTime(cid string, periodTime [3]uint64) {
	t.periodTime[cid] = periodTime
}
