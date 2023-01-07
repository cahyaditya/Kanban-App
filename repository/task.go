package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	con := r.db.WithContext(ctx)
	res := []entity.Task{}
	rows := con.Model(&res).Select("*").Where("user_id = ?", id).Scan(&res)
	if rows.Error != nil {
		return nil, rows.Error
	}
	return res, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	con := r.db.WithContext(ctx)
	temp := entity.Task{}
	err = con.Create(&task).Scan(&temp).Error
	if err != nil {
		return 0, err
	}
	return temp.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	con := r.db.WithContext(ctx)
	res := entity.Task{}
	err := con.Model(&res).Where("id = ?", id).Scan(&res).Error
	if err != nil {
		return entity.Task{}, err
	}
	return res, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	con := r.db.WithContext(ctx)
	res := []entity.Task{}
	rows := con.Model(&res).Select("*").Where("category_id = ?", catId).Find(&res).Scan(&res)
	if rows.Error != nil {
		return nil, rows.Error
	}
	return res, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	con := r.db.WithContext(ctx)
	temp := entity.Task{}
	err := con.Model(&temp).Where("id = ?", task.ID).Updates(&task).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	con := r.db.WithContext(ctx)
	err := con.Where("id = ?", id).Delete(&entity.Task{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
