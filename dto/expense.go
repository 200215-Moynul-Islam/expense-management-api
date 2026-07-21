package dto

import "time"

type CreateExpenseRequest struct {
	CategoryID int     `json:"category_id" valid:"Required"`
	Title      string  `json:"title" valid:"Required;MaxSize(255)"`
	Amount     float64 `json:"amount" valid:"Required"`
	Note       string  `json:"note" valid:"MaxSize(1000)"`
	ExpenseDate string `json:"expense_date" valid:"Required"`
}

type UpdateExpenseRequest struct {
	CategoryID int     `json:"category_id" valid:"Required"`
	Title      string  `json:"title" valid:"Required;MaxSize(255)"`
	Amount     float64 `json:"amount" valid:"Required"`
	Note       string  `json:"note" valid:"MaxSize(1000)"`
	ExpenseDate string `json:"expense_date" valid:"Required"`
}

type GetExpensesRequest struct {
	CategoryID *int
	FromDate   *time.Time
	ToDate     *time.Time

	Page  *int
	Limit *int

	SortBy    string
	SortOrder string
}
