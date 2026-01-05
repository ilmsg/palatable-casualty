package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ilmsg/palatable-casualty/dto"
	"github.com/ilmsg/palatable-casualty/model"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func (h *UserHandler) Find(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	h.db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	json.NewDecoder(r.Body).Decode(&userDto)
	newUser := model.User{Email: userDto.Email, Password: userDto.Password}
	h.db.Create(&newUser)

	var user model.User
	h.db.First(&user, "id = ?", newUser.ID)
	json.NewEncoder(w).Encode(user)
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

type AuthorizeHandler struct {
	db *gorm.DB
}

func (h *AuthorizeHandler) Find(w http.ResponseWriter, r *http.Request) {
	var authorizes []model.Authorize
	h.db.Preload("Actions").Find(&authorizes)
	json.NewEncoder(w).Encode(authorizes)
}

func (h *AuthorizeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var authorizeDto dto.AuthorizeDto
	json.NewDecoder(r.Body).Decode(&authorizeDto)

	newAuthorize := model.Authorize{
		Title:  authorizeDto.Title,
		RoleId: authorizeDto.RoleId,
	}
	h.db.Create(&newAuthorize)

	var actions []*model.Action
	for _, action := range authorizeDto.Actions {
		actions = append(actions, &model.Action{Title: action, AuthorizeId: newAuthorize.ID})
	}
	h.db.Create(&actions)

	var authorize model.Authorize
	h.db.Preload("Actions").First(&authorize, "id = ?", newAuthorize.ID)
	json.NewEncoder(w).Encode(authorize)
}

func NewAuthorize(db *gorm.DB) *AuthorizeHandler {
	return &AuthorizeHandler{db}
}

type RoleHandler struct {
	db *gorm.DB
}

func (h *RoleHandler) Find(w http.ResponseWriter, r *http.Request) {
	var roles []model.Role
	h.db.Preload("Authorizes").Find(&roles)
	json.NewEncoder(w).Encode(roles)
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var roleDto dto.RoleDto
	json.NewDecoder(r.Body).Decode(&roleDto)
	newRole := model.Role{Title: roleDto.Title}
	h.db.Create(&newRole)

	var role model.Role
	h.db.Preload("Authorizes").First(&role, "id = ?", newRole.ID)
	json.NewEncoder(w).Encode(role)
}

func NewRole(db *gorm.DB) *RoleHandler {
	return &RoleHandler{db}
}

type MemberRoleHandler struct {
	db *gorm.DB
}

func (h *MemberRoleHandler) Find(w http.ResponseWriter, r *http.Request) {
	var memberRoles []model.MemberRole
	h.db.Find(&memberRoles)
	json.NewEncoder(w).Encode(memberRoles)
}

func (h *MemberRoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var memberRoleDto dto.MemberRoleDto
	json.NewDecoder(r.Body).Decode(&memberRoleDto)
	newMemberRole := model.MemberRole{
		UserId:    memberRoleDto.UserId,
		RoleId:    memberRoleDto.RoleId,
		ProjectId: memberRoleDto.ProjectId,
	}
	h.db.Create(&newMemberRole)

	var memberRole model.MemberRole
	h.db.First(&memberRole, "id = ?", newMemberRole.ID)
	json.NewEncoder(w).Encode(memberRole)
}

func NewMemberRoleHandler(db *gorm.DB) *MemberRoleHandler {
	return &MemberRoleHandler{db}
}

type ProjectHandler struct {
	db *gorm.DB
}

func (h *ProjectHandler) Find(w http.ResponseWriter, r *http.Request) {
	var projects []model.Project
	h.db.Preload("Members.User").
		Preload("Members.Role").
		Preload("Members.Role.Authorizes").
		Preload("Members.Role.Authorizes.Actions").
		Find(&projects)
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var projectDto dto.ProjectDto
	json.NewDecoder(r.Body).Decode(&projectDto)
	newProject := model.Project{Title: projectDto.Title}
	h.db.Create(&newProject)

	var project model.Project
	h.db.Preload("Members").First(&project, "id = ?", newProject.ID)
	json.NewEncoder(w).Encode(project)
}

func NewProjectHandler(db *gorm.DB) *ProjectHandler {
	return &ProjectHandler{db}
}
