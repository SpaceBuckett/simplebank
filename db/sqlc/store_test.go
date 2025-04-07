package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simplebank/db/utils"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := utils.RandomBankBalance()

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		require.NotEmpty(t, result.Transfer)
		require.Equal(t, result.Transfer.FromAccountID, account1.ID)
		require.Equal(t, result.Transfer.ToAccountID, account2.ID)
		require.Equal(t, result.Transfer.Amount, amount)
		require.NotZero(t, result.Transfer.CreatedAt)
		require.NotZero(t, result.Transfer.ID)

		require.NotEmpty(t, result.FromEntry)
		require.Equal(t, result.FromEntry.AccountID, account1.ID)
		require.Equal(t, result.FromEntry.Amount, -amount)

		require.NotEmpty(t, result.ToEntry)
		require.Equal(t, result.ToEntry.AccountID, account2.ID)
		require.Equal(t, result.ToEntry.Amount, amount)
	}

}
