package book

type RequestBook struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Price       int    `json:"price" form:"price" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Rating      int    `json:"rating" form:"rating" binding:"required"`
	Discount    int    `json:"discount" form:"discount" binding:"required"`
}

type UpdateBook struct {
	Title       string `json:"title" form:"title"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Rating      int    `json:"rating" form:"rating"`
	Discount    int    `json:"discount" form:"discount"`
}
