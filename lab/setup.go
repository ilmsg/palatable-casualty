//go:build ignore

package main

import (
	"github.com/ilmsg/palatable-casualty/database"
	"github.com/ilmsg/palatable-casualty/model"
)

func main() {
	db := database.NewDatabaseWithSqlite("palatable-casualty.db")
	db.AutoMigrate(
		model.User{},
		model.Authorize{},
		model.Role{},
		model.MemberRole{},
		model.Permission{},
		model.RoleAuthorize{},
		model.Project{},
	)

	u1 := model.User{Email: "scott@gmail.com", Password: "scott"}
	db.Create(&u1)
	u2 := model.User{Email: "tiger@gmail.com", Password: "tiger"}
	db.Create(&u2)

	a1 := model.Authorize{Title: "Project"}
	db.Create(a1)
	a2 := model.Authorize{Title: "Task"}
	db.Create(a2)

	p1 := model.Permission{Title: "Create"}
	db.Create(p1)
	p2 := model.Permission{Title: "Read"}
	db.Create(p2)
	p3 := model.Permission{Title: "Update"}
	db.Create(p3)
	p4 := model.Permission{Title: "Delete"}
	db.Create(p4)

	r1 := model.Role{Title: "Project Owner"}
	db.Create(r1)
	r2 := model.Role{Title: "Member"}
	db.Create(r2)

	// Project Owner -> Project, Task ->
	ra1 := model.RoleAuthorize{RoleId: r1.ID, AuthorizeId: a1.ID, Permissions: []model.Permission{p1, p2, p3, p4}}
	ra2 := model.RoleAuthorize{RoleId: r1.ID, AuthorizeId: a2.ID, Permissions: []model.Permission{p1, p2, p3, p4}}
	db.Create(ra1)
	db.Create(ra2)

	// Member -> Project, Task ->
	ra3 := model.RoleAuthorize{RoleId: r2.ID, AuthorizeId: a1.ID, Permissions: []model.Permission{p2}}
	ra4 := model.RoleAuthorize{RoleId: r2.ID, AuthorizeId: a2.ID, Permissions: []model.Permission{p2, p3}}
	db.Create(ra3)
	db.Create(ra4)

	project1 := model.Project{
		Title: "Project 1",
		Members: []model.MemberRole{
			{UserId: u1.ID, RoleId: r1.ID},
			{UserId: u2.ID, RoleId: r2.ID},
		},
	}
	db.Create(&project1)

}
