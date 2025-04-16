package film

type Handler struct {
	usecase Usecase
}

func NewHandler(u Usecase) Handler {
	return Handler{
		usecase: u,
	}
}
