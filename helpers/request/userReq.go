package request

type User struct {
	UserId    int64  `json:"userId,omitempty" form:"userId"`
	FirstName string `json:"firstName" form:"firstName" binding:"required"`
	LastName  string `json:"lastName" form:"lastName" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	MobileNo  string `json:"mobileNo" form:"mobileNo" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	Gender    int    `json:"gender,omitempty" form:"gender"`
}
