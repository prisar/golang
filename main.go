package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

type User struct {
	Name string `json:"Name"`
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Descrition", Content: "Hello World"},
	}
	fmt.Println("All articles endpoint hit")
	json.NewEncoder(w).Encode(articles)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Called Post API of articles")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/articles", postArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":443", router))
}

func main() {

	handleRequests()

	db, err := sql.Open("mysql", "pritimaysaura:daddydidadeadlydeed@tcp(db4free.net:3306)/shopping_schema")

	if err != nil {
		panic(err.Error)
	}
	defer db.Close()

	fmt.Println("Successfully connected to mysql database")

	// insert, err := db.Query("INSERT INTO `shopping_schema`.`user` (`first_name`, `last_name`, `email`, `password`, `phone_number`) VALUES ('test', 'go', 'testgo1@gmail.com', '07a1fe7cfa9c519c78eeed4e099ba603', '9298383831') ")

	// if err != nil {
	// 	panic(err.Error)
	// }
	// defer insert.Close()
	// fmt.Println("Successfully inserted into mysql database")

	results, err := db.Query("Select first_name from user")
	if err != nil {
		panic(err.Error)
	}
	defer results.Close()

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}
}
