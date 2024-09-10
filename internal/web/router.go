package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-pet-project/internal/services"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func getIssuesHandler(w http.ResponseWriter, r *http.Request) {

	issueList := services.GetIssueList()
	js, err := json.Marshal(issueList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func StartRoutes() {

	port := ":8080"

	http.Handle("/", http.FileServer(http.Dir("./web")))
	// TODO what is StripPrefix and the difference ?
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	//http.Handle("/assets/", http.StripPrefix("/", http.FileServer(http.Dir("assets"))))

	//http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/api/v1/issues", getIssuesHandler)

	fmt.Println("Web server started!")
	http.ListenAndServe(port, nil)
}
