package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Transaction struct {
	BactId        string `json:"bactid"`
	EventId       string `json:"eventid"`
	UserId        string `json:"userid"`
	TransactionId string `json:"transactionid"`
}

//var transactions []Transaction

func main() {
	r := mux.NewRouter()
	//routing
	r.HandleFunc("/create-transaction", createTransaction).Methods("POST")
	r.HandleFunc("/", serveHome).Methods("GET")
	//listen to port
	log.Fatal(http.ListenAndServe(":8080", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API <h1>"))
}

func createTransaction(w http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var tran Transaction
	if request.PostFormValue("bactid") == "" || request.PostFormValue("eventid") == "" || request.PostFormValue("userid") == "" || request.PostFormValue("transactionid") == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		tran.BactId = request.PostFormValue("bactid")
		tran.EventId = request.PostFormValue("eventid")
		tran.UserId = request.PostFormValue("userid")
		tran.TransactionId = request.PostFormValue("transactionid")
		//transactions = append(transactions, tran)

		w.Header().Set("content-Type", "application/json")
		json.NewEncoder(w).Encode(tran)

	}

	//Access the photo key - First Approach
	// _, h, err := request.FormFile("photo")
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// fmt.Printf("The type of file : %T", h)
	// saveFile(h, "formfile")
	// //Access the photo key - Second Approach
	// for _, h := range request.MultipartForm.File["photo"] {
	// 	err := saveFile(h, "mapaccess")
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// }
	w.WriteHeader(200)
	return
}
