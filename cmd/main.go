//go:build ignore

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilmsg/palatable-casualty/database"
	"github.com/ilmsg/palatable-casualty/handler"
	"github.com/ilmsg/palatable-casualty/model"
)

func main() {
	db := database.NewDatabaseWithSqlite("database.db")
	db.AutoMigrate(
		model.Project{},
		model.MemberRole{},
		model.Action{},
		model.Authorize{},
		model.Role{},
		model.User{},
	)

	hUser := handler.NewUserHandler(db)
	hProject := handler.NewProjectHandler(db)
	hRole := handler.NewRole(db)
	hAuthorize := handler.NewAuthorize(db)
	hMemberRole := handler.NewMemberRoleHandler(db)

	r := mux.NewRouter()
	rDashboard := r.PathPrefix("/dashboard").Subrouter()

	rUser := rDashboard.PathPrefix("/users").Subrouter()
	rUser.HandleFunc("", hUser.Find).Methods(http.MethodGet)
	rUser.HandleFunc("", hUser.Create).Methods(http.MethodPost)

	rAuthorize := rDashboard.PathPrefix("/authorizes").Subrouter()
	rAuthorize.HandleFunc("", hAuthorize.Find).Methods(http.MethodGet)
	rAuthorize.HandleFunc("", hAuthorize.Create).Methods(http.MethodPost)

	rRole := rDashboard.PathPrefix("/roles").Subrouter()
	rRole.HandleFunc("", hRole.Find).Methods(http.MethodGet)
	rRole.HandleFunc("", hRole.Create).Methods(http.MethodPost)

	rMemberRole := rDashboard.PathPrefix("/members").Subrouter()
	rMemberRole.HandleFunc("", hMemberRole.Find).Methods(http.MethodGet)
	rMemberRole.HandleFunc("", hMemberRole.Create).Methods(http.MethodPost)

	rProject := rDashboard.PathPrefix("/projects").Subrouter()
	rProject.HandleFunc("", hProject.Find).Methods(http.MethodGet)
	rProject.HandleFunc("", hProject.Create).Methods(http.MethodPost)

	log.Println("Server listen at *:7020")
	log.Fatal(http.ListenAndServe(":7020", r))
}
