package main

import(
//	"fmt"
	"net/http"
	"encoding/json"
	"time"
//	"html/template"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)
//render를 사용하여 template를 사용하려면
//꼭 templates폴더안에 xx.tmpl 확장자로 파일이 있어야한다.
var rd *render.Render

type User struct{
	Name string `json: "name"`
	Email string `json : "email"`
	CreatedAt time.Time `json : "createdat"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request){
	user := User{Name : "dongjun", Email : "dongjun@naver.com"}
	rd.JSON(w, http.StatusOK, user)
//	w.Header().Add("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	data, _ := json.Marshal(user)
//	fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request){
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		fmt.Fprint(w,err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)
//	w.Header().Add("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	data, _ := json.Marshal(user)
//	fmt.Fprint(w, string(data))

}

func helloHandler(w http.ResponseWriter, r *http.Request){
//	tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		fmt.Fprint(w,err)
//		rd.Text(w, http.StatusBadRequest, err.Error())
//		return
//	}
	user := User{Name : "dongjun", Email : "dongjun@naver.com"}
	rd.HTML(w, http.StatusOK, "body", user)
	//tmpl.ExecuteTemplate(w, "hello.tmpl", "DongJun")

}

func main(){
//render를 사용하여 template를 사용하려면
//꼭 templates폴더안에 xx.tmpl 확장자로 파일이 있어야한다.
//하지만 render.Option에 Extensions키 값에 인식할 수 있는 확장자명을 추가할 수 있다. 그리고 폴더명은 Directory키 값에 추가할 수 있다.
	rd = render.New(render.Options{
		Directory : "template",
		Extensions:[]string{".html",".tmpl"},
		Layout : "hello",
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
