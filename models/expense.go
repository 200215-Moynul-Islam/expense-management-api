package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Expense struct {
	ID          int        `orm:"column(id);auto;pk" json:"id"`
	User        *User      `orm:"column(user_id);rel(fk)" json:"user,omitempty"`
	Category    *Category  `orm:"column(category_id);rel(fk)" json:"category,omitempty"`
	Title       string     `orm:"column(title);size(255)" json:"title"`
	Amount      float64    `orm:"column(amount);digits(12);decimals(2)" json:"amount"`
	Note        string     `orm:"column(note);null;type(text)" json:"note,omitempty"`
	ExpenseDate time.Time  `orm:"column(expense_date);type(date)" json:"expense_date"`
	CreatedAt   time.Time  `orm:"column(created_at);auto_now_add;type(timestamp)" json:"created_at"`
	UpdatedAt   time.Time  `orm:"column(updated_at);auto_now;type(timestamp)" json:"updated_at"`
}

func (e *Expense) TableName() string {
	return "expenses"
}

func init() {
	orm.RegisterModel(new(Expense))
}
