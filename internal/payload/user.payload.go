package payload

type CreateUser struct {
	Email      string `json:"email" validate:"email,required"`
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	ProfileUrl string `json:"profileUrl" validate:"required"`
}
