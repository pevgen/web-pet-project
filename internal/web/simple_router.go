package web

import (
	"fmt"
	"net/http"
	"web-pet-project/internal/services"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func getIssuesHandler(w http.ResponseWriter, r *http.Request) {

	fileBytes, err := services.GetIssueListAsCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(fileBytes)
}

func StartRoutes() {

	port := ":8080"

	http.Handle("/", http.FileServer(http.Dir("./web")))
	// TODO what is StripPrefix and the difference ?
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/api/v1/issues", getIssuesHandler)

	fmt.Println("Web server started!")
	http.ListenAndServe(port, nil)
}
