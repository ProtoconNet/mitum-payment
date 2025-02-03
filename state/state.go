package state

import (
	"fmt"
	"strings"

	"github.com/ProtoconNet/mitum-payment/types"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/pkg/errors"
)

var (
	DesignStateValueHint  = hint.MustNewHint("mitum-payment-design-state-value-v0.0.1")
	PaymentStateKeyPrefix = "payment"
	DesignStateKeySuffix  = "design"
)

func PaymentStateKey(addr string) string {
	return fmt.Sprintf("%s:%s", PaymentStateKeyPrefix, addr)
}

type DesignStateValue struct {
	hint.BaseHinter
	Design types.Design
}

func NewDesignStateValue(design types.Design) DesignStateValue {
	return DesignStateValue{
		BaseHinter: hint.NewBaseHinter(DesignStateValueHint),
		Design:     design,
	}
}

func (sv DesignStateValue) Hint() hint.Hint {
	return sv.BaseHinter.Hint()
}

func (sv DesignStateValue) IsValid([]byte) error {
	e := util.ErrInvalid.Errorf("invalid DesignStateValue")

	if err := sv.BaseHinter.IsValid(DesignStateValueHint.Type().Bytes()); err != nil {
		return e.Wrap(err)
	}

	if err := sv.Design.IsValid(nil); err != nil {
		return e.Wrap(err)
	}

	return nil
}

func (sv DesignStateValue) HashBytes() []byte {
	return sv.Design.Bytes()
}

func GetDesignFromState(st base.State) (types.Design, error) {
	v := st.Value()
	if v == nil {
		return types.Design{}, errors.Errorf("state value is nil")
	}

	d, ok := v.(DesignStateValue)
	if !ok {
		return types.Design{}, errors.Errorf("expected DesignStateValue but %T", v)
	}

	return d.Design, nil
}

func IsDesignStateKey(key string) bool {
	return strings.HasPrefix(key, PaymentStateKeyPrefix) && strings.HasSuffix(key, DesignStateKeySuffix)
}

func DesignStateKey(addr string) string {
	return fmt.Sprintf("%s:%s", PaymentStateKey(addr), DesignStateKeySuffix)
}

var (
	AccountRecordStateValueHint = hint.MustNewHint("mitum-payment-account-record-state-value-v0.0.1")
	AccountRecordStateKeySuffix = "accountrecord"
)

type AccountRecordStateValue struct {
	hint.BaseHinter
	AccountRecord types.AccountRecord
}

func NewAccountRecordStateValue(accountRecord types.AccountRecord) AccountRecordStateValue {
	return AccountRecordStateValue{
		BaseHinter:    hint.NewBaseHinter(AccountRecordStateValueHint),
		AccountRecord: accountRecord,
	}
}

func (sv AccountRecordStateValue) Hint() hint.Hint {
	return sv.BaseHinter.Hint()
}

func (sv AccountRecordStateValue) IsValid([]byte) error {
	e := util.ErrInvalid.Errorf("invalid TimeStampLastIdxStateValue")

	if err := sv.BaseHinter.IsValid(AccountRecordStateValueHint.Type().Bytes()); err != nil {
		return e.Wrap(err)
	}

	return nil
}

func (sv AccountRecordStateValue) HashBytes() []byte {
	return util.ConcatBytesSlice(sv.AccountRecord.Bytes())
}

func GetAccountRecordFromState(st base.State) (*types.AccountRecord, error) {
	v := st.Value()
	if v == nil {
		return nil, errors.Errorf("state value is nil")
	}

	isv, ok := v.(AccountRecordStateValue)
	if !ok {
		return nil, errors.Errorf("expected AccountRecordStateValue but, %T", v)
	}

	return &isv.AccountRecord, nil
}

func IsAccountRecordStateKey(key string) bool {
	return strings.HasPrefix(key, PaymentStateKeyPrefix) && strings.HasSuffix(key, AccountRecordStateKeySuffix)
}

func AccountRecordStateKey(addr string, acAddr string) string {
	return fmt.Sprintf("%s:%s:%s", PaymentStateKey(addr), acAddr, AccountRecordStateKeySuffix)
}

//var (
//	ItemStateValueHint = hint.MustNewHint("mitum-payment-item-state-value-v0.0.1")
//	ItemStateKeySuffix = "item"
//)
//
//type ItemStateValue struct {
//	hint.BaseHinter
//	Item types.AccountInfo
//}
//
//func NewItemStateValue(item types.AccountInfo) ItemStateValue {
//	return ItemStateValue{
//		BaseHinter: hint.NewBaseHinter(ItemStateValueHint),
//		Item:       item,
//	}
//}
//
//func (sv ItemStateValue) Hint() hint.Hint {
//	return sv.BaseHinter.Hint()
//}
//
//func (sv ItemStateValue) IsValid([]byte) error {
//	e := util.ErrInvalid.Errorf("invalid ItemStateValue")
//
//	if err := sv.BaseHinter.IsValid(ItemStateValueHint.Type().Bytes()); err != nil {
//		return e.Wrap(err)
//	}
//
//	if err := sv.Item.IsValid(nil); err != nil {
//		return e.Wrap(err)
//	}
//
//	return nil
//}
//
//func (sv ItemStateValue) HashBytes() []byte {
//	return sv.Item.Bytes()
//}
//
//func GetItemFromState(st base.State) (types.AccountInfo, error) {
//	v := st.Value()
//	if v == nil {
//		return types.AccountInfo{}, errors.Errorf("State value is nil")
//	}
//
//	ts, ok := v.(ItemStateValue)
//	if !ok {
//		return types.AccountInfo{}, common.ErrTypeMismatch.Wrap(errors.Errorf("expected ItemStateValue found, %T", v))
//	}
//
//	return ts.Item, nil
//}
//
//func IsItemStateKey(key string) bool {
//	return strings.HasPrefix(key, PaymentStateKeyPrefix) && strings.HasSuffix(key, ItemStateKeySuffix)
//}
//
//func ItemStateKey(addr base.Address, pid string, index uint64) string {
//	return fmt.Sprintf("%s:%s:%s:%s", TimeStampStateKey(addr), pid, strconv.FormatUint(index, 10), ItemStateKeySuffix)
//}
