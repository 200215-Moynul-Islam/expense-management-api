package repositories

import (
	"expense-management-api/models"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ExpenseFilter struct {
	CategoryID *int

	FromDate *time.Time
	ToDate   *time.Time

	Page  *int
	Limit *int

	SortBy    string
	SortOrder string
}

type ExpenseRepository interface {
	Create(expense *models.Expense) error
	GetExpenses(
		userID int,
		filter ExpenseFilter,
	) ([]*models.Expense, error)

	GetByID(
		id int,
	) (*models.Expense, error)

	Delete(
		expense *models.Expense,
	) error

	Update(expense *models.Expense) error
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

func (r *expenseRepository) GetExpenses(
	userID int,
	filter ExpenseFilter,
) ([]*models.Expense, error) {

	o := orm.NewOrm()

	query := o.QueryTable(new(models.Expense)).
	Filter("user_id", userID).
	RelatedSel()

	if filter.CategoryID != nil {
		query = query.Filter(
			"category_id",
			*filter.CategoryID,
		)
	}

	if filter.FromDate != nil {
		query = query.Filter(
			"expense_date__gte",
			*filter.FromDate,
		)
	}

	if filter.ToDate != nil {
		query = query.Filter(
			"expense_date__lte",
			*filter.ToDate,
		)
	}

	if filter.SortBy != "" {

		sortField := filter.SortBy

		if filter.SortOrder == "desc" {
			sortField = "-" + sortField
		}

		query = query.OrderBy(sortField)
	}

	if filter.Page != nil && filter.Limit != nil {

		offset := (*filter.Page - 1) * (*filter.Limit)

		query = query.Offset(offset)
		query = query.Limit(*filter.Limit)
	}

	var expenses []*models.Expense

	_, err := query.All(&expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *expenseRepository) GetByID(
	id int,
) (*models.Expense, error) {

	o := orm.NewOrm()

	var expense models.Expense

	err := o.QueryTable(new(models.Expense)).
		Filter("id", id).
		RelatedSel().
		One(&expense)

	if err == orm.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func (r *expenseRepository) Delete(
	expense *models.Expense,
) error {

	o := orm.NewOrm()

	_, err := o.Delete(expense)

	return err
}

func (r *expenseRepository) Update(
	expense *models.Expense,
) error {

	o := orm.NewOrm()

	_, err := o.Update(expense)

	return err
}
