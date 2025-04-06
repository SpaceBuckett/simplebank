package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simplebank/db/utils"
	"testing"
	"time"
)

func createRandomTransfer(t *testing.T, account1 Account, account2 Account, amount int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createRandomTransfer(t, account1, account2, utils.RandomBankBalance())
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createdTransfer := createRandomTransfer(t, account1, account2, utils.RandomBankBalance())

	receivedTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)

	require.NoError(t, err)
	require.Equal(t, createdTransfer.ID, receivedTransfer.ID)
	require.Equal(t, createdTransfer.FromAccountID, receivedTransfer.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, receivedTransfer.ToAccountID)
	require.Equal(t, createdTransfer.Amount, receivedTransfer.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt, receivedTransfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2, utils.RandomBankBalance())
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)
}
