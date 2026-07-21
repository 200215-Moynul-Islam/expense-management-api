package repositories

import (
	"expense-management-api/models"

	"github.com/beego/beego/v2/client/orm"
)

type ExpenseRepository interface {
	Create(expense *models.Expense) error
}

type expenseRepository struct{}

func NewExpenseRepository() ExpenseRepository {
	return &expenseRepository{}
}

func (r *expenseRepository) Create(
	expense *models.Expense,
) error {

	o := orm.NewOrm()

	_, err := o.Insert(expense)

	return err
}
