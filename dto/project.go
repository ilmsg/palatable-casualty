package dto

type ProjectDto struct {
	Title string `json:"title"`
}

type UserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RoleDto struct {
	Title string `json:"title"`
}

type AuthorizeDto struct {
	Title   string   `json:"title"`
	RoleId  uint     `json:"role_id"`
	Actions []string `json:"actions"`
}

type MemberRoleDto struct {
	UserId    uint `json:"user_id"`
	RoleId    uint `json:"role_id"`
	ProjectId uint `json:"project_id"`
}
