package dto

type CreateExpenseRequest struct {
	CategoryID int     `json:"category_id" valid:"Required"`
	Title      string  `json:"title" valid:"Required;MaxSize(255)"`
	Amount     float64 `json:"amount" valid:"Required"`
	Note       string  `json:"note" valid:"MaxSize(1000)"`
	ExpenseDate string `json:"expense_date" valid:"Required"`
}
