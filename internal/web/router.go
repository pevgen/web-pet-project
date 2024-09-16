package web

import (
	"fmt"
	"github.com/gocraft/web"
	"log"
	"net/http"
	"path"
	"web-pet-project/internal/dbms/repository/mongodb"
	"web-pet-project/internal/services"
)

type Context struct {
	HelloCount int
}

// var issueService = services.NewIssuesService(memory.NewIssuesRepository())
// var issueService = services.NewIssuesService(postgres.NewIssueRepository())
var issueService = services.NewIssuesService(mongodb.NewIssuesRepository())

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *Context) csvFileFromIssuesHandler(w web.ResponseWriter, r *web.Request) {

	// Handlers should read before writing
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=file1")
	// should set up after headers
	w.WriteHeader(http.StatusOK)

	bytes, err := issueService.GetIssueListAsCsv()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (c *Context) getIssuesHandler(w web.ResponseWriter, r *web.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, err := issueService.GetIssueListAsJson()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)

}
func StartRoutesWithLib() {

	port := ":8080"

	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware).     // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
					Middleware(web.StaticMiddleware(path.Join("./web"))).
					Middleware(web.StaticMiddleware(path.Join("./assets"))).
					Middleware((*Context).SetHelloCount). // Your own middleware!
					Get("/api/v1/issues", (*Context).getIssuesHandler).
					Get("/api/v1/issues/files/csv", (*Context).csvFileFromIssuesHandler)

	fmt.Println("Web server started!")

	log.Fatal(http.ListenAndServe(port, router))
}
