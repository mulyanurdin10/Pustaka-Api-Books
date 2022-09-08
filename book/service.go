package book

type Service interface {
	FindAll() ([]Books, error)
	FindById(id int) (Books, error)
	Create(requestBook RequestBook) (Books, error)
	Update(id int, updateBook UpdateBook) (Books, error)
	Delete(id int) (Books, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Books, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(id int) (Books, error) {
	return s.repository.FindById(id)
}

func (s *service) Create(requestBook RequestBook) (Books, error) {

	book := Books{
		Title:       requestBook.Title,
		Price:       requestBook.Price,
		Discount:    requestBook.Discount,
		Rating:      requestBook.Rating,
		Description: requestBook.Description,
	}

	return s.repository.Create(book)
}

func (s *service) Update(id int, updateBooks UpdateBook) (Books, error) {
	book, _ := s.repository.FindById(id)

	book.Title = updateBooks.Title
	book.Price = updateBooks.Price
	book.Discount = updateBooks.Discount
	book.Rating = updateBooks.Rating
	book.Description = updateBooks.Description

	updateBook, err := s.repository.Update(book)

	return updateBook, err
}

func (s *service) Delete(id int) (Books, error) {
	book, _ := s.repository.FindById(id)
	deleteBook, err := s.repository.Delete(book)
	return deleteBook, err
}
