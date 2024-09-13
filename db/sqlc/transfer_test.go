package db

import (
	"context"
	"testing"

	"github.com/logeshwarann-dev/gophers-bank-app/util"
	"github.com/stretchr/testify/require"
)

func createOnlyTransfer(t *testing.T, arg CreateTransferParams) Transfer {

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer

}

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoneyWithinLimit(account1.Balance),
	}

	transfer := createOnlyTransfer(t, arg)
	//	require.Equal(t, account1.Balance, account1.Balance-transfer.Amount)
	//	require.Equal(t, account2.Balance, account2.Balance+transfer.Amount)
	return transfer

}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	transfer1, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer1)
	require.Equal(t, transfer.ID, transfer1.ID)
	require.Equal(t, transfer.FromAccountID, transfer1.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer1.ToAccountID)
	require.Equal(t, transfer.Amount, transfer1.Amount)
}

func TestListTransfers(t *testing.T) {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 4; i++ {
		arg := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        util.RandomMoneyWithinLimit(account1.Balance),
		}
		createOnlyTransfer(t, arg)
	}

	transferArg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Offset:        2,
		Limit:         2,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), transferArg)
	require.NoError(t, err)
	require.Len(t, transfers, 2)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}
