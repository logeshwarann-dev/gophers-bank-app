package db

import (
	"context"
	"testing"
	"time"

	"github.com/logeshwarann-dev/gophers-bank-app/util"
	"github.com/stretchr/testify/require"
)

func createOnlyEntry(t *testing.T, arg CreateEntryParams) Entry {

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func createRandomEntry(t *testing.T) Entry {
	account1 := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}

	entry := createOnlyEntry(t, arg)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {

	account1 := createRandomAccount(t)

	for i := 0; i < 4; i++ {
		arg := CreateEntryParams{
			AccountID: account1.ID,
			Amount:    util.RandomMoney(),
		}
		createOnlyEntry(t, arg)
	}

	entryArg := ListEntriesParams{
		AccountID: account1.ID,
		Offset:    2,
		Limit:     2,
	}

	entries, err := testQueries.ListEntries(context.Background(), entryArg)
	require.NoError(t, err)
	require.Len(t, entries, 2)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
