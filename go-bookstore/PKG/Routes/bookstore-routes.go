package routes

import(
	"github.com/jaywantadh/go-bookstore/PKG/Controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", Controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", Controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", Controllers.UpdateBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", Controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", Controllers.DeleteBook).Methods("DELETE")
}

