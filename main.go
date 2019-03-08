package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/satriaa14/test-driven-deployment/models"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3031"
	}

	svr := ":" + port

	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/hasil", routeSubmitPost)

	fmt.Printf("server started at localhost%s", svr)
	http.ListenAndServe(svr, nil)
}

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("./src/index.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("./src/index.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var num1 = r.Form.Get("num1")
		num1Conv, _ := strconv.ParseFloat(num1, 64)

		var num2 = r.Form.Get("num2")
		num2Conv, _ := strconv.ParseFloat(num2, 64)
		var tipe = r.Form.Get("tipe")

		var nilai = models.Nilai{A: num1Conv, B: num2Conv, Tipe: tipe}

		result, _ := nilai.Perhitungan()

		resultStr := fmt.Sprintf("%f", result)

		var data = map[string]string{"num1": num1,
			"num2":   num2,
			"tipe":   tipe,
			"result": resultStr}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
