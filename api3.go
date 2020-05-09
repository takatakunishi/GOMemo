package main

import (
	"encoding/json"
    // "fmt"
    // "io/ioutil"
    // "log"
	"net/http"
	// "net/http/cgi"

    // "github.com/julienschmidt/httprouter"
)

//ref = https://konboi.hatenablog.com/entry/2014/09/23/172756

type Ping struct {
	Status int `json:"status"`
	Result string `json:"result"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, "ok"}

	res, err := json.Marshal(ping)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main (){
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":3001", nil)
}