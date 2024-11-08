package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Sandhya-Pratama/simple-bank/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := sqlc.NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	fmt.Println(">> before:", account1.Balance, account2.Balance)
	// run n concurrent transfer transactions
	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), sqlc.TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()

	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}
	// existed := make(map[int]bool)
	// for i := 0; i < n; i++ {
	// 	err := <-errs
	// 	require.NoError(t, err)

	// 	result := <-results
	// 	require.NotEmpty(t, result)

	// 	// check transfer
	// 	transfer := result.Transfer
	// 	require.NotEmpty(t, transfer)
	// 	require.Equal(t, account1.ID, transfer.FromAccountID)
	// 	require.Equal(t, account2.ID, transfer.ToAccountID)
	// 	require.Equal(t, amount, transfer.Amount)
	// 	require.NotZero(t, transfer.ID)
	// 	require.NotZero(t, transfer.CreatedAt)

	// 	_, err = store.GetTransfer(context.Background(), transfer.ID)
	// 	require.NoError(t, err)

	// 	// check entries
	// 	fromEntry := result.FromEntry
	// 	require.NotEmpty(t, fromEntry)
	// 	require.Equal(t, account1.ID, fromEntry.AccountID)
	// 	require.Equal(t, -amount, fromEntry.Amount)
	// 	require.NotZero(t, fromEntry.ID)
	// 	require.NotZero(t, fromEntry.CreatedAt)

	// 	_, err = store.GetEntrie(context.Background(), fromEntry.ID)
	// 	require.NoError(t, err)

	// 	toEntry := result.ToEntry
	// 	require.NotEmpty(t, toEntry)
	// 	require.Equal(t, account2.ID, toEntry.AccountID)
	// 	require.Equal(t, amount, toEntry.Amount)
	// 	require.NotZero(t, toEntry.ID)
	// 	require.NotZero(t, toEntry.CreatedAt)

	// 	_, err = store.GetEntrie(context.Background(), toEntry.ID)
	// 	require.NoError(t, err)

	// 	// check account
	// 	fromAccount := result.FromAccount
	// 	require.NotEmpty(t, fromAccount)
	// 	require.Equal(t, account1.ID, fromAccount.ID)

	// 	toAccount := result.ToAccount
	// 	require.NotEmpty(t, toAccount)
	// 	require.Equal(t, account2.ID, toAccount.ID)

	// 	//check account's balance
	// 	diff1 := account1.Balance - fromAccount.Balance
	// 	diff2 := toAccount.Balance - account2.Balance
	// 	require.Equal(t, diff1, diff2)
	// 	require.True(t, diff1 > 0)
	// 	require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ...

	// 	k := int(diff1 / amount)
	// 	require.True(t, k >= 1 && k <= n)
	// 	require.NotContains(t, existed, k)
	// 	existed[k] = true
	// }

	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after :", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}
