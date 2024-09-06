package accounts

type Repository interface {
	Save(a Account) error
	Retrieve(accountID int64) (Account, error)
	RetrieveAll() ([]Account, error)
}
