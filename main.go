package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/tidwall/pretty"
)

type Response struct {
	Result string `json:"result"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		jsonStr := r.Form.Get("json")

		var jsonObj interface{}
		err := json.Unmarshal([]byte(jsonStr), &jsonObj)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		formattedJson := pretty.Pretty([]byte(jsonStr))

		response := Response{Result: string(formattedJson)}
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
