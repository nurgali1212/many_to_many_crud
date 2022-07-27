package service

import (
	"fmt"
	"rest_api_golang_crud_sqlx/model"
	"rest_api_golang_crud_sqlx/repository"
)

type BookListServive interface {
	GetAll() ([]model.Book, error)
	Create(book model.Book) (string, error)
	GetById(id string) (model.Book, error)
	Delete(id string) error
	Update(input model.Book) error
}
type CreateBookInput struct {
	Title      string `json:"title"`
	Author     string `json:"author" `

}

type CreateCategoryinput struct {
	Genre	string	`json:"genre"`
}

type CreateBookCategoryinput struct {
	Book_id     uint `json:"book_id" db:"book_id"`
	Category_id uint `json:"category_id" db:"category_id"`

}
type CategoryService interface {
	GetAllCategoryService() ([]model.Category, error)
	CreateC(category model.Category) (string, error)
	GetByIdC(id string) (model.Category, error)
	DeleteC(id string) error
	UpdateC(input model.Category) error
}

type BookCategoryService interface {
	GetAllBookCategoryService() ([]model.BookCategory, error)
	CreateBookCategoryService(bookcategory model.BookCategory) (string, error)
	GetByIdBookCategoryService(id string) (model.BookCategory, error)
	DeleteBookCategoryService(id string) error
	UpdateBookCategoryService(input model.BookCategory) error

}
type Service struct {
	repo *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		repo: repos,
	}
}
// get Category
func (s *Service) GetAllCategoryService() ([]model.Category, error) {
	return s.repo.GetAllCategoryRepo()
}
//Create Category
func (s *Service) CreateC(input model.Category) (model.Category, error) {
	return s.repo.CreateCategoryRepo(input)
}
//get id Category
func (s *Service) GetByIdC(id string) (model.Category, error) {
	return s.repo.GetByIdCategoryRepo(id)
}
// put Category
func (s *Service) UpdateC(categoryId int, input model.UpdateCategoryinput) error {
	category, err := s.repo.GetByIdCategoryRepo(fmt.Sprint(categoryId))
	if err != nil {
		return err
	}
	category.Genre = *input.Genre
	
	return s.repo.UpdateCategory(category)
}
//deleteCategory
func (s *Service) DeleteC(id string) error {
	return s.repo.DeleteCategoryRepo(id)
}
//GET BOOK
func (s *Service) GetAll() ([]model.Book, error) {
	return s.repo.GetAllBook()
}
//GET ID BOOK
func (s *Service) GetById(id string) (model.Book, error) {
	return s.repo.GetByIdBook(id)
}

//PUT BOOK
func (s *Service) Update(bookId int, input model.UpdateBookInput) error {
	book, err := s.repo.GetByIdBook(fmt.Sprint(bookId))
	if err != nil {
		return err
	}
	book.Author = *input.Author
	book.Title = *input.Title
	return s.repo.UpdateBook(book)
}
//DELETE BOOK
func (s *Service) Delete(id string) error {
	return s.repo.DeleteBook(id)
}
// POST BOOK
func (s *Service) Create(input model.Book) (model.Book, error) {

	return s.repo.CreateBook(input)
}



// get BookCategory
func (s *Service) GetAllBookCategoryService() ([]model.BookCategory, error) {
	return s.repo.GetAllBookCategoryRepo()
}


//Create Category
func (s *Service) CreateBookCategoryService(input model.BookCategory) (model.BookCategory, error) {
	return s.repo.CreateBookCategoryRepo(input)
}
//get id Category
func (s *Service) GetByIdBookCategoryService(id string) (model.BookCategory, error) {
	return s.repo.GetByIdBookCategoryRepo(id)
}
// put Category
func (s *Service) UpdateBookCategoryService(bookcategoryId int, input model.UpdateBookCategoryinput) error {
	bookcategory, err := s.repo.GetByIdBookCategoryRepo(fmt.Sprint(bookcategoryId))
	if err != nil {
		return err
	}
	bookcategory.Book_id = *input.Book_id
	bookcategory.Category_id = *input.Category_id
	
	return s.repo.UpdateBookCategoryRepo(bookcategory)
}
//deleteCategory
func (s *Service) DeleteBookCategoryService(id string) error {
	return s.repo.DeleteBookCategoryRepo(id)
}