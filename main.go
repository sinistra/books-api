package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"

	"sinistra/books-api/controllers"
	"sinistra/books-api/driver"
	"sinistra/books-api/models"

	"github.com/gorilla/mux"
)

var books []models.Book
var db *sql.DB
var port string

func init() {
	godotenv.Load()
	var ok bool
	port, ok = os.LookupEnv("HOST_PORT")
	if !ok {
		port = "8000"
	}
}

func main() {
	//log.Println("Port="+port)
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		panic("oh no")
	})

	n := negroni.New()

	recovery := negroni.NewRecovery()
	recovery.Formatter = &negroni.HTMLPanicFormatter{}

	n.Use(recovery)
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	fmt.Println("Server is running at port " + port)
	address := ":" + port
	log.Fatal(http.ListenAndServe(address, n))
}
