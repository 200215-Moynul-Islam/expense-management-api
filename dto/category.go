package dto

type CreateCategoryRequest struct {
	Name string `json:"name" valid:"Required;MaxSize(100)"`
}
