package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	con := r.db.WithContext(ctx)
	res := []entity.Category{}
	rows := con.Model(&res).Select("*").Where("user_id = ?", id).Scan(&res)
	if rows.Error != nil {
		return nil, rows.Error
	}
	return res, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	con := r.db.WithContext(ctx)
	temp := entity.Category{}
	err = con.Create(&category).Scan(&temp).Error
	if err != nil {
		return 0, err
	}
	return temp.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	con := r.db.WithContext(ctx)
	err := con.Create(&categories).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	con := r.db.WithContext(ctx)
	res := entity.Category{}
	err := con.Model(&res).Where("id = ?", id).Scan(&res).Error
	if err != nil {
		return entity.Category{}, err
	}
	return res, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	con := r.db.WithContext(ctx)
	temp := entity.Category{}
	err := con.Model(&temp).Where("id = ?", category.ID).Updates(&category).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	con := r.db.WithContext(ctx)
	err := con.Where("id = ?", id).Delete(&entity.Category{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
