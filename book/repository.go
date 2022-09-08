package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Books, error)
	FindById(id int) (Books, error)
	Create(books Books) (Books, error)
	Update(books Books) (Books, error)
	Delete(books Books) (Books, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Books, error) {
	var books []Books
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindById(id int) (Books, error) {
	var book Books
	err := r.db.First(&book, id).Error
	return book, err
}

func (r *repository) Create(book Books) (Books, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(book Books) (Books, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Delete(book Books) (Books, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
