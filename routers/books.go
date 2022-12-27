package routers

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"api.com/my_api/Database"

	//"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ToAbsPath(fileName string) string {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		panic(err)
	}

	return absPath
}

func Handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// Serve the resource.
		id := req.FormValue("id")
		println(id)
		if id == "" {
			D := DataBase.Records.ToJSON(w)
			w.Write(D)
			return
		}
		intId, _ := strconv.Atoi(id)
		if intId-1 >= len(DataBase.Records.Records) {
			println(true)
			http.Error(w, "No records Found at id "+string(id), http.StatusNotFound)
			return
		}
		println(len(DataBase.Records.Records))
		D := DataBase.Records.ToJSONId(w, intId-1)
		w.Write(D)
	case http.MethodPost:
		var newRecord DataBase.Books
		newRecord.Tobook(w, req.Body)
		DataBase.Records.Records = append(DataBase.Records.Records, newRecord)
		fmt.Println(DataBase.Records.Records)
		err := ioutil.WriteFile(ToAbsPath("./database/db.json"), DataBase.Records.ToJSON(w), fs.ModeAppend)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Data added Successfully"))
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
