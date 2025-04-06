package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simplebank/db/utils"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T, account Account) Enterei {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomBankBalance(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	createdEntry := createRandomEntry(t, account)

	receivedEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)

	require.NoError(t, err)
	require.Equal(t, createdEntry.ID, receivedEntry.ID)
	require.Equal(t, createdEntry.Amount, receivedEntry.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt, receivedEntry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    2,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)
}
