package user

type Usecase struct {
	repository Repository
}

func NewUsecase() Usecase {
	return Usecase{}
}
