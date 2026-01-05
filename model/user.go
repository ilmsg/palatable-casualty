package model

type ActionType string

const (
	ActionCreate ActionType = "Create"
	ActionRead   ActionType = "Read"
	ActionUpdate ActionType = "Update"
	ActionDelete ActionType = "Delete"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Action struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"` // Create, Read, Update, Delete
	AuthorizeId uint
}

type Authorize struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"` // Project, Task
	Actions []Action
	RoleId  uint
}

type Role struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"` // ProjectOwner, Member
	Authorizes []Authorize
}

type MemberRole struct {
	ID        uint `json:"id"`
	UserId    uint `json:"-"`
	User      User `gorm:"foreignKey:UserId"`
	RoleId    uint `json:"-"`
	Role      Role `gorm:"foreignKey:RoleId"`
	ProjectId uint
}

type Project struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Members []MemberRole
}
