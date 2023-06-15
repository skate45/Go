package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID 		string `json: "id"`
	Title 	string `json: "title"`
	Author 	string `json: "author"`
	Quantity uint  `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

//curl localhost:8080/book
func getBooks(context *gin.Context){
	context.IndentedJSON(http.StatusOK,books)
}

func getBookById(id string)(*book, error){
	for i,bookItem:=range books{
		if bookItem.ID==id{
			return &books[i],nil
		}
	}
	return nil, errors.New("Book not found")
}

//curl localhost:8080/book/3
func findBookById(context *gin.Context){
	id:=context.Param("id")
	book,err:=getBookById(id)
	
	if(err!=nil){
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found"})
		return
	}
	context.IndentedJSON(http.StatusOK,book)
}

//curl localhost:8080/checkout?id=2 --request "PATCH"
func checkoutBook(context *gin.Context){
	id,ok:=context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Missing id query parameter"})
		return
	}

	book, err:=getBookById(id)
	if err!=nil {
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found"})
		return
	}
	if book.Quantity<=0 {
		context.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Book not available"})
		return
	}

	book.Quantity-=1
	context.IndentedJSON(http.StatusOK,book)
}

//curl localhost:8080/return?id=2 --request "PATCH"
func returnBook(context *gin.Context){
	id,ok:=context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Missing id query parameter"})
		return
	}

	book, err:=getBookById(id)
	if err!=nil {
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found"})
		return
	}

	book.Quantity+=1
	context.IndentedJSON(http.StatusOK,book)
}
//curl localhost:8080/book --include --header "Context-type: application/json" -d @body.json --request "POST"
func createBook(context *gin.Context){
	var newBook book
	if err:= context.BindJSON(&newBook); err!=nil {
		return
	}
	books=append(books,newBook)
	context.IndentedJSON(http.StatusCreated,newBook)
}

func main(){
	router:=gin.Default()

	router.GET("/book",getBooks)

	router.GET("/book/:id",findBookById)

	router.POST("/book",createBook)

	router.PATCH("/checkout",checkoutBook)

	router.PATCH("/return",returnBook)

	router.Run("localhost:8080")

}