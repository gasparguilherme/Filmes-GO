package film

type Usecase struct {
	repository Repository
}

func NewUsecase(r Repository) Usecase {
	concreteUsecase := Usecase{
		repository: r,
	}
	return concreteUsecase
}
