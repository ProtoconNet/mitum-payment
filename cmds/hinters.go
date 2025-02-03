package cmds

import (
	currencycmds "github.com/ProtoconNet/mitum-currency/v3/cmds"
	"github.com/ProtoconNet/mitum-payment/operation/payment"
	"github.com/ProtoconNet/mitum-payment/state"
	"github.com/ProtoconNet/mitum-payment/types"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/pkg/errors"
)

var Hinters []encoder.DecodeDetail
var SupportedProposalOperationFactHinters []encoder.DecodeDetail

var AddedHinters = []encoder.DecodeDetail{
	// revive:disable-next-line:line-length-limit

	{Hint: types.DesignHint, Instance: types.Design{}},
	{Hint: types.AccountInfoHint, Instance: types.AccountInfo{}},
	{Hint: types.AccountRecordHint, Instance: types.AccountRecord{}},

	{Hint: payment.DepositHint, Instance: payment.Deposit{}},
	{Hint: payment.RegisterModelHint, Instance: payment.RegisterModel{}},
	{Hint: payment.TransferHint, Instance: payment.Transfer{}},
	{Hint: payment.UpdateAccountInfoHint, Instance: payment.UpdateAccountInfo{}},

	{Hint: state.DesignStateValueHint, Instance: state.DesignStateValue{}},
	{Hint: state.AccountRecordStateValueHint, Instance: state.AccountRecordStateValue{}},
}

var AddedSupportedHinters = []encoder.DecodeDetail{
	{Hint: payment.DepositFactHint, Instance: payment.DepositFact{}},
	{Hint: payment.RegisterModelFactHint, Instance: payment.RegisterModelFact{}},
	{Hint: payment.TransferFactHint, Instance: payment.TransferFact{}},
	{Hint: payment.UpdateAccountInfoFactHint, Instance: payment.UpdateAccountInfoFact{}},
}

func init() {
	Hinters = append(Hinters, currencycmds.Hinters...)
	Hinters = append(Hinters, AddedHinters...)

	SupportedProposalOperationFactHinters = append(SupportedProposalOperationFactHinters, currencycmds.SupportedProposalOperationFactHinters...)
	SupportedProposalOperationFactHinters = append(SupportedProposalOperationFactHinters, AddedSupportedHinters...)
}

func LoadHinters(encs *encoder.Encoders) error {
	for i := range Hinters {
		if err := encs.AddDetail(Hinters[i]); err != nil {
			return errors.Wrap(err, "add hinter to encoder")
		}
	}

	for i := range SupportedProposalOperationFactHinters {
		if err := encs.AddDetail(SupportedProposalOperationFactHinters[i]); err != nil {
			return errors.Wrap(err, "add supported proposal operation fact hinter to encoder")
		}
	}

	return nil
}
