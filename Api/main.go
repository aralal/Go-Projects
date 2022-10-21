package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Transaction struct {
	BactId        string `json:"bactid"`
	EventId       string `json:"eventid"`
	UserId        string `json:"userid"`
	TransactionId string `json:"transactionid"`
	Status        bool   `json:"status"`
}

//var transactions []Transaction

func main() {
	r := mux.NewRouter()
	fmt.Println("API started")
	//routing
	r.HandleFunc("/create-transaction", createTransaction).Methods("POST")
	r.HandleFunc("/home", serveHome).Methods("GET")
	//listen to port
	log.Fatal(http.ListenAndServe(":8090", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serve function invoked")
	w.Write([]byte("<h1>Welcome to the API <h1>"))
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create transaction invoked")
	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var tran Transaction
	if r.PostFormValue("bactid") == "" || r.PostFormValue("eventid") == "" || r.PostFormValue("userid") == "" || r.PostFormValue("transactionid") == "" {

		tran.Status = false
		tran.BactId = r.PostFormValue("bactid")
		tran.EventId = r.PostFormValue("eventid")
		tran.UserId = r.PostFormValue("userid")
		tran.TransactionId = r.PostFormValue("transactionid")
		//w.Header().Set("content-Type", "application/json")
		//json.NewEncoder(w).Encode(tran)
		//file, err := os.Create("./trans.json")
		file, _ := json.MarshalIndent(tran, "", " ")
		ipfsfile = ioutil.WriteFile("./trans.json", file, 0644)

	} else {
		// tran.Status = true
		// tran.BactId = r.PostFormValue("bactid")
		// tran.EventId = r.PostFormValue("eventid")
		// tran.UserId = r.PostFormValue("userid")
		// tran.TransactionId = r.PostFormValue("transactionid")
		// w.Header().Set("content-Type", "application/json")
		// json.NewEncoder(w).Encode(tran)

	}

	w.WriteHeader(200)
	return
}
