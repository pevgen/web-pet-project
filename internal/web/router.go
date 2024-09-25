package web

import (
	"fmt"
	"github.com/gocraft/web"
	"net/http"
	"path"
	"strconv"
	"web-pet-project/internal/config"
	"web-pet-project/internal/dbms/repository"
	"web-pet-project/internal/dbms/repository/memory"
	"web-pet-project/internal/dbms/repository/mongodb"
	"web-pet-project/internal/dbms/repository/postgres"
	"web-pet-project/internal/services"
)

type Context struct {
	//HelloCount int
}

var issueService services.IssuesService

// var issueService = services.NewIssuesService(postgres.NewIssueRepository())
//var issueService = services.NewIssuesService(mongodb.NewIssuesRepository())

//func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
//	c.HelloCount = 3
//	next(rw, req)
//}

func NewRouter(cfg config.AppConfig) (func(), error) {

	issueService = services.NewIssuesService(
		[]repository.IssuesRepository{
			memory.NewIssuesRepository(),
			postgres.NewIssueRepository(cfg.Db.Postgres.ConnectString),
			mongodb.NewIssuesRepository(cfg.Db.Mongodb.ConnectString, cfg.Db.Mongodb.DbName),
		})

	port := ":" + strconv.Itoa(cfg.WebServer.Port)
	router := setupRouter()

	fmt.Println("Web server started!")
	err := http.ListenAndServe(port, router)

	cancel := func() {
		issueService.CloseRepos()
	}
	if err != nil {
		return cancel, err
	}

	return cancel, nil
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

func setupRouter() *web.Router {

	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware).     // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
					Middleware(web.StaticMiddleware(path.Join("./web"))).
					Middleware(web.StaticMiddleware(path.Join("./assets"))).
		//Middleware((*Context).SetHelloCount). // Your own middleware!
		Get("/api/v1/issues", (*Context).getIssuesHandler).
		Get("/api/v1/issues/files/csv", (*Context).csvFileFromIssuesHandler)
	return router
}
