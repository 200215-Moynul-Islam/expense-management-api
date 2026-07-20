package repositories

import (
	"expense-management-api/models"

	"github.com/beego/beego/v2/client/orm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetByNameAndUserID(name string, userID int) (*models.Category, error)
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) Create(
	category *models.Category,
) error {

	o := orm.NewOrm()

	_, err := o.Insert(category)

	return err
}

func (r *categoryRepository) GetByNameAndUserID(
	name string,
	userID int,
) (*models.Category, error) {

	o := orm.NewOrm()

	category := &models.Category{}

	err := o.QueryTable(
		new(models.Category),
	).
		Filter("name", name).
		Filter("user_id", userID).
		One(category)

	if err == orm.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return category, nil
}
