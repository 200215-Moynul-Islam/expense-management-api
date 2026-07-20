package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Category struct {
	ID        int         `orm:"column(id);auto;pk" json:"id"`
	User      *User       `orm:"column(user_id);rel(fk)" json:"user,omitempty"`
	Name      string      `orm:"column(name);size(100)" json:"name"`
	CreatedAt time.Time   `orm:"column(created_at);auto_now_add;type(timestamp)" json:"created_at"`
	UpdatedAt time.Time   `orm:"column(updated_at);auto_now;type(timestamp)" json:"updated_at"`
	Expenses  []*Expense  `orm:"reverse(many)" json:"expenses,omitempty"`
}

func (c *Category) TableName() string {
	return "categories"
}

func init() {
	orm.RegisterModel(new(Category))
}
