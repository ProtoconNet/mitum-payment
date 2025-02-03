package cmds

type PaymentCommand struct {
	Deposit           DepositCommand           `cmd:"" name:"deposit" help:"deposit"`
	UpdateAccountInfo UpdateAccountInfoCommand `cmd:"" name:"update-account-info" help:"update account info"`
	Transfer          TransferCommand          `cmd:"" name:"transfer" help:"transfer"`
	RegisterModel     RegisterModelCommand     `cmd:"" name:"register-model" help:"register payment model"`
}
