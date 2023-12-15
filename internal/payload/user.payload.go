package payload

type Register struct {
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required"`
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	ProfileUrl string `json:"profileUrl" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUser struct {
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	Phone      string `json:"phone" validate:"min=10"`
	ProfileUrl string `json:"profileUrl" validate:"required"`
}

type ChangePassword struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
