package usecase

type ListUsecase interface {
	GetLists() (ips []string, err error)
}

type listUsecaseImpl struct {
}

func NewListUsecase() ListUsecase {
	return &listUsecaseImpl{}
}

func (u *listUsecaseImpl) GetLists() (ips []string, err error) {
	return []string{""}, nil
}
