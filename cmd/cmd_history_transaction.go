package cmd

import (
	"context"
	"fmt"

	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/transactions"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/platform"
)

type CmdHistoryTransaction struct {
}

func (c CmdHistoryTransaction) Execute(ctx context.Context) error {
	service := ctx.Value(platform.TransactionServiceKey)

	var transactionService transactions.TransactionService = service.(transactions.TransactionService)
	transfer, err := transactionService.TransferList()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < len(transfer); i++ {
		fmt.Printf("History of transfer %v %.2f", transfer[i].Type, transfer[i].Ammount)
	}
	return nil
}
