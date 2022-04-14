package task_aggregate

import (
	"context"
	"mygofsm/my_order_project/domain/task_domain/task_aggregate/entity"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

type TaskRepositoryImpl struct {
}

func (impl *TaskRepositoryImpl) CreateTask(ctx *context.Context, taskEntity *entity.Task) error {
	return nil
}
func (impl *TaskRepositoryImpl) GetTaskById(ctx *context.Context, taskId int64) (*entity.Task, error) {
	return nil, nil
}
func (impl *TaskRepositoryImpl) GetTaskResourceIds(ctx *context.Context, taskId int64) (*entity.Resource, error) {
	return nil, nil
}
