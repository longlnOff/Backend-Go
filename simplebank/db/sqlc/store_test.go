package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// writing database transaction is something we must always be very careful with.
// It can be easy to write, but can also easily become a nightmare if we don’t handle the concurrency carefully.
// So the best way to make sure that our transaction works well is to run it with several concurrent go routines.
// -> We should test the function with concurrent go routine
func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	n := 50
	amount := int64(10)

	// run n concurrent transfer transaction

	// 1. create a channel to store the error
	errs := make(chan error)
	// results := make(chan TransferTxResult)
	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID
		if i % 2 == 0 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		txName := fmt.Sprintf("tx %d", i + 1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			_, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID: toAccountID,
				Amount: amount,
			})

			errs <- err
			// results <- result

		} ()
	}

	// check results
	// existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <- errs
		require.NoError(t, err)
	}
	// 	result := <- results
	// 	require.NotEmpty(t, result)

	// 	// check transfer
	// 	transfer := result.Transfer
	// 	require.NotEmpty(t, transfer)
	// 	require.Equal(t, account1.ID, transfer.FromAccountID.Int64)
	// 	require.Equal(t, account2.ID, transfer.ToAccountID.Int64)
	// 	require.Equal(t, amount, transfer.Amount)
	// 	require.NotZero(t, transfer.ID)
	// 	require.NotZero(t, transfer.CreatedAt)

	// 	_, err = store.GetTransfer(context.Background(), transfer.ID)
	// 	require.NoError(t, err)

	// 	// check accounts
	// 	fromAccount := result.FromAccount
	// 	require.NotEmpty(t, fromAccount)
	// 	require.Equal(t, account1.ID, fromAccount.ID)

	// 	toAccount := result.ToAccount
	// 	require.NotEmpty(t, toAccount)
	// 	require.Equal(t, account2.ID, toAccount.ID)

	// 	// check account's balance
	// 	fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
	// 	diff1 := account1.Balance - fromAccount.Balance
	// 	diff2 := toAccount.Balance - account2.Balance
	// 	require.Equal(t, diff1, diff2)
	// 	require.True(t, diff1 > 0)
	// 	require.True(t, diff1 % amount == 0)

	// 	k := int(diff1 / amount)
	// 	require.True(t, k >= 1 && k <= n)

	// 	require.NotContains(t, existed, k)
	// 	existed[k] = true
	// }

	// check the final updated balance
	updateAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updateAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	// fmt.Println((">> After transfer: "), updateAccount1.Balance, updateAccount2.Balance)
	require.Equal(t, account1.Balance, updateAccount1.Balance)
	require.Equal(t, account2.Balance, updateAccount2.Balance)

	// require.Equal(t, account1.Balance - int64(n) * amount, updateAccount1.Balance)
	// require.Equal(t, account2.Balance + int64(n) * amount, updateAccount2.Balance)
}