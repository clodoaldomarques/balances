package accounts

type Service struct {
	rep Repository
}

func New(r Repository) *Service {
	return &Service{
		rep: r,
	}
}

func (s Service) CreateNewAccount(a Account) error {

	return nil
}

func (s Service) UpdateExistentAccount(a Account) error {
	return nil
}
