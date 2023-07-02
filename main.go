package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

var indexHTML = `
<!DOCTYPE html>
<html>
<body>
	<h1>JSON Formatter</h1>
	<form action="/format" method="post">
		<textarea name="json" rows="10" cols="30"></textarea>
		<br/>
		<input type="submit" value="Format">
	</form>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index").Parse(indexHTML)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
	
		inputJSON := r.FormValue("json")
	
		var v interface{}
		err := json.Unmarshal([]byte(inputJSON), &v)
	
		// If there's an error, return the unformatted JSON
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(inputJSON))
			return
		}
	
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})
	
	http.ListenAndServe(":8080", nil)
}
