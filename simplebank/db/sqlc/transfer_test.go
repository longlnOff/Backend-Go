package db


func createRandomTransfer(t *testing.T) {
	arg := CreateTransferParams{
		FromAccountID: util.Ra,
	}
}


func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)