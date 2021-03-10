package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"github.com/gorilla/pat"
)

type User struct{
	Name string `json: "name"`
	Email string `json : "email"`
	CreatedAt time.Time `json : "createdat"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request){

}

func addUserHandler(w http.ResponseWriter, r *httpRequest){

}

func helloHandler(w http.ResponseWriter, r *http.Request){

}

func main(){
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	http.ListenAndServe(":3000", mux)
}
