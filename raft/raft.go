package raft

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

type cacheManager struct {
	data map[string]string
	sync.RWMutex
}


func Get(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	key := params["key"]
	fmt.Println(key)
}

func Set(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["key"]
	value := params["value"]
	fmt.Println(key,value)
}

func HttpServe() {
	r := mux.NewRouter()
	r.HandleFunc("/get/{key}", Get)
	r.HandleFunc("/set/{key}/{value}", Set)
	log.Fatal(http.ListenAndServe(":8080", r))
}