package dto

type CreateCategoryRequest struct {
	Name string `json:"name" valid:"Required;MaxSize(100)"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" valid:"Required;MaxSize(100)"`
}
