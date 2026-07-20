package dto

type RegisterRequest struct {
	Name     string `json:"name" valid:"Required;MaxSize(100)"`
	Email    string `json:"email" valid:"Required;Email;MaxSize(255)"`
	Password string `json:"password" valid:"Required;MinSize(6);MaxSize(255)"`
}

type LoginRequest struct {
	Email    string `json:"email" valid:"Required;Email;MaxSize(255)"`
	Password string `json:"password" valid:"Required;MinSize(6);MaxSize(255)"`
}
