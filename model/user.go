package model

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Authorize struct {
	ID    uint   `json:"id"`
	Title string `json:"title"` // Project, Task
}

type Permission struct {
	ID    uint   `json:"id"`
	Title string `json:"title"` // Create, Read, Update, Delete
}

type Role struct {
	ID    uint   `json:"id"`
	Title string `json:"title"` // ProjectOwner, Member
}

type RoleAuthorize struct {
	ID          uint `json:"id"`
	RoleId      uint
	AuthorizeId uint
	Permissions []Permission
}

type MemberRole struct {
	ID     uint `json:"id"`
	UserId uint
	RoleId uint
}

type Project struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Members []MemberRole
}
