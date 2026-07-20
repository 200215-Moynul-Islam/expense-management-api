package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	ID         int          `orm:"column(id);auto;pk" json:"id"`
	Name       string       `orm:"column(name);size(100)" json:"name"`
	Email      string       `orm:"column(email);size(255);unique" json:"email"`
	Password   string       `orm:"column(password);size(255)" json:"-"`
	CreatedAt  time.Time    `orm:"column(created_at);auto_now_add;type(timestamp)" json:"created_at"`
	UpdatedAt  time.Time    `orm:"column(updated_at);auto_now;type(timestamp)" json:"updated_at"`
	Categories []*Category  `orm:"reverse(many)" json:"categories,omitempty"`
	Expenses   []*Expense   `orm:"reverse(many)" json:"expenses,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(User))
}
