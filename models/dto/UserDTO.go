package dto

type UserDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required,min=2,max=30"`
	Email    string `json:"email" validate:"required,customEmail,max=50"`
	Password string `json:"password" validate:"required,min=5,max=12"`
}
