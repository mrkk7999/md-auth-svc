package usergroup

type UserToGroupRequest struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	OldGroup string `json:"oldGroup"`
	NewGroup string `json:"newGroup"`
}
