package film

type Usecase struct {
	repository Repository
}

func NewUsecase(r Repository) Usecase {
	usecaseConcreto := Usecase{
		repository: r,
	}
	return usecaseConcreto
}
