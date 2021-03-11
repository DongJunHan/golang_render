package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"html/template"
	"github.com/gorilla/pat"
)

type User struct{
	Name string `json: "name"`
	Email string `json : "email"`
	CreatedAt time.Time `json : "createdat"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request){
	user := User{Name : "dongjun", Email : "dongjun@naver.com"}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request){
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,err)
		return
	}
	user.CreatedAt = time.Now()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))

}

func helloHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w,err)
		return
	}
	tmpl.ExecuteTemplate(w, "hello.tmpl", "DongJun")

}

func main(){
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	http.ListenAndServe(":3000", mux)
}
