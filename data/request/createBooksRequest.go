package request

type CreateBooksRequest struct {
	BookName string `validate:"required,min=1,max=200" json:"name"`
}
