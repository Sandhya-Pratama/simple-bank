package test

import (
	"context"
	"testing"

	"github.com/Sandhya-Pratama/simple-bank/db/sqlc"
	"github.com/Sandhya-Pratama/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
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

}
