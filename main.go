package main

import (
	"encoding/json"
	"io/ioutil"

	DataBase "api.com/my_api/Database"
	"api.com/my_api/routers"

	//"api.com/my_api/routes"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fileContent, err := ioutil.ReadFile("E:/Learnings/golang/api/database/db.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(fileContent)
}

func main() {
	data, err := ioutil.ReadFile("./database/db.json")
	if err != nil {
		panic(err)
	}
	err2 := json.Unmarshal(data, &DataBase.Records)
	if err != nil {
		panic(err2)
	}
	print("\n")
	fmt.Print(DataBase.Records)
	print("\n")
	fmt.Println("Listening on  http://localhost:8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/books", routers.Handler)
	fmt.Println(DataBase.Records)
	http.ListenAndServe(":8080", mux)

}
