package repository

import (
	"context"
	"mygofsm/my_order_project/domain/task_domain/task_aggregate/entity"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

type TaskRepository interface {
	CreateTask(ctx *context.Context, taskEntity *entity.Task) error
	GetTaskById(ctx *context.Context, taskId int64) (*entity.Task, error)
	GetTaskResourceIds(ctx *context.Context, taskId int64) (*entity.Resource, error)
}
