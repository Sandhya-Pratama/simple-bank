package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Sandhya-Pratama/simple-bank/db/sqlc"
	"github.com/Sandhya-Pratama/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) sqlc.Account {
	args := sqlc.CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// test create account
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// test get account
func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// test update account
func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	// Mengupdate balance dengan nilai acak
	newBalance := util.RandomMoney()
	arg := sqlc.UpdateAccountParams{
		ID:      account1.ID,
		Balance: newBalance,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	// Verifikasi bahwa pembaruan balance sesuai
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, newBalance, account2.Balance) // Perbarui baris ini
	require.Equal(t, account1.Currency, account2.Currency)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// test delete account
func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	// Hapus akun
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	// Coba dapatkan akun yang dihapus
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	// Pastikan error yang didapat adalah sql.ErrNoRows
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, account2)
}

// test list account
func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := sqlc.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
