package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	con := r.db.WithContext(ctx)
	res := entity.User{}
	err := con.Model(&res).Where("id = ?", id).Scan(&res).Error
	if err != nil {
		return entity.User{}, err
	}
	return res, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	con := r.db.WithContext(ctx)
	res := entity.User{}
	err := con.Model(&res).Where("email = ?", email).Scan(&res).Error
	if err != nil {
		return entity.User{}, err
	}
	return res, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	con := r.db.WithContext(ctx)
	temp := entity.User{}
	err := con.Create(&user).Scan(&temp)
	if err.Error != nil {
		return entity.User{}, err.Error
	}
	return temp, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	con := r.db.WithContext(ctx)
	temp := entity.User{}
	err := con.Model(&temp).Where("id = ?", user.ID).Updates(&user).Scan(&temp).Error
	if err != nil {
		return entity.User{}, err
	}
	return temp, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	con := r.db.WithContext(ctx)
	err := con.Where("id = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
