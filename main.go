package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"rsc.io/quote/v3"
)

var gitCommit string

func main() {
	fmt.Println("starting http server ")
	r := mux.NewRouter()
	r.HandleFunc("/", helloworld)
	r.HandleFunc("/go", goquote)
	r.HandleFunc("/opt", opttruth)
        r.HandleFunc("/version",version)

	s := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	 fmt.Fprintf(w,"Hello World")
         fmt.Fprintf(w, quote.HelloV3())
}

func goquote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, quote.GoV3())
}

func opttruth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, quote.GoV3())
}

func version(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w,"myapplication: ["+
"  {"+
"version: 1.0,\n"+
"lastcommitsha: %s,\n"+
"description : pre-interview technical test\n"+
"} ]",gitCommit)
}

