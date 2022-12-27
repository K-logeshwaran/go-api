package DataBase

import (
	"encoding/json"
	"io"
	"net/http"
)

type Books struct {
	Name      string     `json:"name"`
	Ratings   float64    `json:"ratings"`
	Author    string     `json:"author"`
	Price     float64    `json:"price"`
	Publisher Publishers `json:"publisher"`
}

type Publishers struct {
	Name        string `json:"name"`
	Contact     string `json:"contact"`
	No_of_books int    `json:"no_of_books"`
}

type DB struct {
	Records []Books
}

// var Database = []Books{
// 	{Name: "Automate", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
// 	{Name: "Golang", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
// 	{Name: "Node js", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
// 	{Name: "Rust", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
// 	{Name: "Bot", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
// }

var Records = DB{Records: []Books{
	{Name: "Automate", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
	{Name: "Golang", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
	{Name: "Node js", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
	{Name: "Rust", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
	{Name: "Bot", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
}}

func (d DB) ToJSON(w http.ResponseWriter) []byte {
	jsonData, err := json.Marshal(d.Records)
	if err != nil {
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return []byte("Internal server Error")
	}
	return jsonData
}

func (d DB) ToJSONId(w http.ResponseWriter, id int) []byte {
	jsonData, err := json.Marshal(d.Records[id])
	if err != nil {
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return []byte("Internal server Error")
	}
	return jsonData
}

func (b *Books) Tobook(w http.ResponseWriter, body io.ReadCloser) {
	err := json.NewDecoder(body).Decode(b)
	if err != nil {
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
	}
}
