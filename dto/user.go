package dto

type UpdateUserRequest struct {
	Name string `json:"name" valid:"Required;MaxSize(100)"`
}
