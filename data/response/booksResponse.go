package response

type BooksResponse struct {
	BookId   int    `json:"bookId"`
	BookName string `json:"bookName"`
	Author   string `json:"author"`
}
