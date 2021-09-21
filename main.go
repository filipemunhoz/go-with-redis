package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type Person struct {
	Name string `json:"name"`
}

func main() {

	r := mux.NewRouter()
	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/save", save)

	gett := r.Methods(http.MethodGet).Subrouter()
	gett.HandleFunc("/", getAll)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getAll(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	val, err := rdb.Get("person").Result()
	if err != nil {
		fmt.Println(err)
	}

	p := &Person{}
	p.Name = val
	json.NewEncoder(rw).Encode(p)
	jData, _ := json.Marshal(val)
	rw.Write(jData)
}

func save(rw http.ResponseWriter, r *http.Request) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var p Person
	json.Unmarshal(body, &p)

	err = rdb.Set("person", p.Name, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Name)
}
