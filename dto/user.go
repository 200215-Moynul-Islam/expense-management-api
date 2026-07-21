package dto

type UpdateUserRequest struct {
	Name  string `json:"name" valid:"Required;MaxSize(100)"`
	Email string `json:"email" valid:"Required;Email;MaxSize(255)"`
}
