package request

type UpdateBooksRequest struct {
	BookId   int    `validate:"required"`
	BookName string `validate:"required,max=200,min=1" json:"name"`
	Author   string `validate:"required,max=200,min=1" json:"author"`
}
