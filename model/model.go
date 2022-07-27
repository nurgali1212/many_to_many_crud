package model

type Book struct {
	Id     uint   `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}

type Category struct {
	Id    uint   `json:"id"`
	Genre string `json:"genre"`
}

type BookCategory struct {
	Id          uint `json:"id" db:"id"`
	Book_id     uint `json:"book_id" db:"book_id"`
	Category_id uint `json:"category_id" db:"category_id"`
}

type UpdateBookInput struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
}
type UpdateCategoryinput struct {
	Genre *string `json:"genre"`
}
type UpdateBookCategoryinput struct {
	Book_id     *uint `json:"book_id" db:"book_id"`
	Category_id *uint `json:"category_id" db:"category_id"`
}


