package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	shell "github.com/ipfs/go-ipfs-api"
)

type Transaction struct {
	BactId        string `json:"bactid"`
	EventId       string `json:"eventid"`
	UserId        string `json:"userid"`
	TransactionId string `json:"transactionid"`
	Status        bool   `json:"status"`
}

func main() {
	r := mux.NewRouter()
	fmt.Println("API started")
	//routing
	r.HandleFunc("/create-transaction", createTransaction).Methods("POST")
	r.HandleFunc("/home", serveHome).Methods("GET")
	//listen to port
	log.Fatal(http.ListenAndServe(":8000", r))

}

// check for local host connection
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serve function invoked")
	w.Write([]byte("<h1>Welcome to the API <h1>"))
}

func addFile(sh *shell.Shell, text string) (string, error) {
	return sh.Add(strings.NewReader(text))
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create transaction invoked")
	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var tran Transaction
	sh := shell.NewShell("localhost:5001")
	if r.PostFormValue("bactid") == "" || r.PostFormValue("eventid") == "" || r.PostFormValue("userid") == "" || r.PostFormValue("transactionid") == "" {

		tran.Status = false
		tran.BactId = r.PostFormValue("bactid")
		tran.EventId = r.PostFormValue("eventid")
		tran.UserId = r.PostFormValue("userid")
		tran.TransactionId = r.PostFormValue("transactionid")
		//writing object to a json file
		file, _ := json.MarshalIndent(tran, "", " ")
		err := ioutil.WriteFile("./trans.json", file, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//ipfs
		fmt.Println("Adding file to IPFS")
		//reading from the file
		data, err := ioutil.ReadFile("./trans.json")
		if err != nil {
			log.Fatal(err)
		}

		cid, err := addFile(sh, string(data))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("File added with CID:", cid)
		json.NewEncoder(w).Encode(cid)

	} else {
		tran.Status = true
		tran.BactId = r.PostFormValue("bactid")
		tran.EventId = r.PostFormValue("eventid")
		tran.UserId = r.PostFormValue("userid")
		tran.TransactionId = r.PostFormValue("transactionid")
		//writing object to a json file
		file, _ := json.MarshalIndent(tran, "", " ")
		err := ioutil.WriteFile("./trans.json", file, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//ipfs
		fmt.Println("Adding file to IPFS")
		//reading from the file
		data, err := ioutil.ReadFile("./trans.json")
		if err != nil {
			log.Fatal(err)
		}
		cid, err := addFile(sh, string(data))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("File added with CID:", cid)
		json.NewEncoder(w).Encode(cid)

	}

	w.WriteHeader(200)
	return
}
