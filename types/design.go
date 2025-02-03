package types

import (
	"encoding/json"
	"github.com/ProtoconNet/mitum-currency/v3/common"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/ProtoconNet/mitum2/util/valuehash"
	"github.com/pkg/errors"
)

var DesignHint = hint.MustNewHint("mitum-payment-design-v0.0.1")

var maxAccounts = 1000

type Design struct {
	hint.BaseHinter
	accounts map[string]AccountInfo
}

func NewDesign() Design {
	accounts := make(map[string]AccountInfo)
	return Design{
		BaseHinter: hint.NewBaseHinter(DesignHint),
		accounts:   accounts,
	}
}

func (de Design) IsValid([]byte) error {
	if err := util.CheckIsValiders(nil, false,
		de.BaseHinter,
	); err != nil {
		return err
	}

	return nil
}

func (de Design) Bytes() []byte {
	var bac []byte
	if de.accounts != nil {
		ac, _ := json.Marshal(de.accounts)
		bac = valuehash.NewSHA256(ac).Bytes()
	} else {
		bac = []byte{}
	}

	return util.ConcatBytesSlice(bac)
}

func (de Design) Hash() util.Hash {
	return de.GenerateHash()
}

func (de Design) GenerateHash() util.Hash {
	return valuehash.NewSHA256(de.Bytes())
}

func (de Design) Accounts() map[string]AccountInfo {
	return de.accounts
}

func (de Design) Account(account string) *AccountInfo {
	v, found := de.accounts[account]

	if !found {
		return nil
	}

	return &v
}

func (de *Design) AddAccount(account AccountInfo) error {
	de.accounts[account.Address().String()] = account

	if len(de.accounts) > maxAccounts {
		return common.ErrValOOR.Wrap(
			errors.Errorf("accounts over allowed, %d > %d", len(de.accounts), maxAccounts))
	}

	return nil
}

func (de *Design) UpdateAccount(account AccountInfo) error {
	_, found := de.accounts[account.Address().String()]
	if !found {
		return common.ErrValueInvalid.Wrap(
			errors.Errorf("account, %v not registered in service", account.Address()))
	}
	de.accounts[account.Address().String()] = account

	return nil
}

func (de *Design) RemoveAccount(account AccountInfo) error {
	_, found := de.accounts[account.Address().String()]
	if !found {
		return errors.Errorf("account, %v not registered in service", account.Address())
	}

	delete(de.accounts, account.Address().String())

	return nil
}
