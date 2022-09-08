package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connected Error")
	}

	db.AutoMigrate(&book.Books{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.FindAll)
	v1.GET("/books/:id", bookHandler.FindById)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.PUT("/books/:id", bookHandler.Update)
	v1.DELETE("/books/:id", bookHandler.Delete)

	router.Run()

	// dbrepository := book.NewRepository(db)

	// book := book.Book{
	// 	Title:       "One Piece",
	// 	Price:       200000,
	// 	Discount:    10,
	// 	Rating:      5,
	// 	Description: "Film yang digemari oleh semua orang",
	// }

	// dbrepository.Create(book)

	// book, _ := dbrepository.FindById(1)

	// fmt.Println(book)

	// books, _ := dbrepository.FindAll()

	// listBook := []interface{}{}

	// for _, v := range books {
	// 	listBook = append(listBook, v)
	// }

	// fmt.Println(listBook...)

	// book := book.Book{}
	// book.Title = "One Pice"
	// book.Price = 200000
	// book.Discount = 20
	// book.Rating = 10
	// book.Description = "Film yang digemari semua orang dari kalangan muda dan tua"

	// errBookCreate := db.Create(&book).Error

	// if errBookCreate != nil {
	// 	fmt.Println("Error creating book record")
	// }

	// var book book.Book

	// errBookGet := db.First(&book, 2).Error

	// if errBookGet != nil {
	// 	fmt.Println("Error get book")
	// }

	// fmt.Println(book)

	// var book []book.Book

	// errGetAllBook := db.Find(&book).Error

	// if errGetAllBook != nil {
	// 	fmt.Println("Error get all book")
	// }

	// listBook := []interface{}{}

	// for _, v := range book {
	// 	listBook = append(listBook, v)
	// }

	// fmt.Println(listBook...)

	// var book []book.Book

	// errGetAllBook := db.Where("rating = ?", 5).Find(&book).Error

	// if errGetAllBook != nil {
	// 	fmt.Println("Error get all book")
	// }

	// listBook := []interface{}{}

	// for _, v := range book {
	// 	listBook = append(listBook, v)
	// }

	// fmt.Println(listBook...)

	// var book book.Book

	// errGetBook := db.Where("title = ?", "Naruto").First(&book).Error

	// if errGetBook != nil {
	// 	fmt.Println("Error get book")
	// }

	// book.Title = "Captain Tsubatsa"

	// errSaveBook := db.Save(&book).Error

	// if errSaveBook != nil {
	// 	fmt.Println("Error updated book")
	// }

	// var book book.Book

	// errDeleteBook := db.Delete(&book, 2).Error

	// if errDeleteBook != nil {
	// 	fmt.Println("Error deleted book")
	// }

}
