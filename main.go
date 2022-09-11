package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Title  string `json:"title"`
	Pages  int    `json:"pages"`
	Author Author `json:"author"`
	Id     int    `json:"id"`
}

type Author struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
}

func find(books []Book, id int) Book {
	var book = Book{}
	for index, value := range books {
		if value.Id == id {
			book = books[index]
		}
	}
	return book
}

func main() {
	var router = mux.NewRouter()

	var books = []Book{
		{"The Man Without Pasta", 298, Author{"James", "Milkan"}, 75},
		{"Sword Destiny: II", 309, Author{"Leslie", "Peter"}, 62},
		{"Understanding Hailwind Effects", 408, Author{"Dennis", "Wilkinson"}, 98},
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get started by going to '/books'"))
	})

	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		var data, _ = json.Marshal(books)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		var id = mux.Vars(r)["id"]

		var intId, _ = strconv.Atoi(id)
		var book = find(books, intId)

		var data, _ = json.Marshal(book)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	http.ListenAndServe("localhost:5200", router)
}
