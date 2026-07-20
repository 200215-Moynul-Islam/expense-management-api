package dto

type RegisterRequest struct {
	Name     string `json:"name" valid:"Required;MaxSize(100)"`
	Email    string `json:"email" valid:"Required;Email"`
	Password string `json:"password" valid:"Required;MinSize(6)"`
}
