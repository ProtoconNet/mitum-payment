package types

import (
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/ProtoconNet/mitum2/util/hint"
)

func (t *AccountRecord) unpack(
	enc encoder.Encoder,
	ht hint.Hint,
	addr string,
) error {
	t.BaseHinter = hint.NewBaseHinter(ht)
	address, err := base.DecodeAddress(addr, enc)
	if err != nil {
		return err
	}
	t.address = address

	return nil
}
