package plugin

type groupRequest struct {
	UserID string `json:"userID" form:"userID"`
	Group  string `json:"group" form:"group"`
}
