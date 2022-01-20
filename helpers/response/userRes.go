package response

type UserEdit struct {
	UserId    int64  `json:"userId,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type User struct {
	UserId    int64  `json:"userId,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	MobileNo  string `json:"mobileNo,omitempty"`
	Email     string `json:"email,omitempty"`
	Gender    int    `json:"gender,omitempty"`
}