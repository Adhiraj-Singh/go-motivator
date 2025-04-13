package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// just a list of quotes

var quotes = []string{
	"Believe in yourself!",
	"Push through the hard days.",
	"Every step counts.",
	"Keep learning. Keep building.",
	"You're doing better than you think.",
}

func main() {
	// this helps serve the static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// this is the route with the html file
	// there is a better way to do this but for now lets serve it here
	http.HandleFunc("/", serveHome)

	http.HandleFunc("/quote", serveQuote)

	// lets start server now
	http.ListenAndServe(":8080", nil)

}

// serving functions
func serveHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, nil)
}

func serveQuote(w http.ResponseWriter, r *http.Request) {
	// randomly call the quotes
	rand.Seed(time.Now().UnixNano()) // research nix nano
	random := quotes[rand.Intn(len(quotes))]
	json.NewEncoder(w).Encode(map[string]string{"quote": random})
}
