package response

type UserEdit struct {
	UserId    int64  `json:"userId,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}