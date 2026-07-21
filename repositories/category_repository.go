package repositories

import (
	"expense-management-api/models"

	"github.com/beego/beego/v2/client/orm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetByNameAndUserID(name string, userID int) (*models.Category, error)
	GetByID(id int) (*models.Category, error)
	GetAllByUserID(userID int) ([]*models.Category, error)
	Update(category *models.Category) error
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

func (r *categoryRepository) GetByID(
	id int,
) (*models.Category, error) {

	o := orm.NewOrm()

	category := &models.Category{
		ID: id,
	}

	err := o.Read(category)

	if err == orm.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) GetAllByUserID(
	userID int,
) ([]*models.Category, error) {

	o := orm.NewOrm()

	var categories []*models.Category

	_, err := o.QueryTable(new(models.Category)).
		Filter("user_id", userID).
		All(&categories)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) Update(
	category *models.Category,
) error {

	o := orm.NewOrm()

	_, err := o.Update(category)

	return err
}
