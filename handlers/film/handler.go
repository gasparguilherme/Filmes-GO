package film

type Handler struct {
	usecase Usecase
}

func NewHandler() Handler {
	return Handler{}
}
