package rest

import (
	"html/template"
	"log"
	"net/http"
)

func CreateRoutes(r *http.ServeMux) {
	r.HandleFunc("/newuser", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/testforms.html")
	if err != nil {
		log.Println("error open user`s creating form")
	}
	t.Execute(w, nil)

	if r.Method == "POST" {
		log.Println(r.Method)
		// функция для парсинга формы, добавление id и проверка формы + загрузка в бд если всё ок
	}

}
