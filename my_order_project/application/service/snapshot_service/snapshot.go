package snapshot_service

import (
	"context"
	"mygofsm/my_order_project/domain/snapshot_domain/snapshot_aggregate/entity"
	"mygofsm/my_order_project/idl/gen-go/snapshot"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

type ApplicationService interface {
	GenerateSnapshotEntity(ctx context.Context, req snapshot.PreCreateOrderReq) (*entity.TaskDtoSubmitSnapshot, error)
	CheckSnapshotTask(ctx context.Context, snapshotEntity *entity.TaskDtoSubmitSnapshot) error
	CreateTask(ctx *context.Context, snapshotId int64) error
}

type ApplicationServiceImpl struct {
}

func (impl *ApplicationServiceImpl) GenerateSnapshotEntity(ctx context.Context, req snapshot.PreCreateOrderReq) (*entity.TaskDtoSubmitSnapshot, error) {
	snapshotEntity := &entity.TaskDtoSubmitSnapshot{
		Id:            0,
		UserId:        0,
		SubmitType:    0,
		SubmitCommand: "",
		ReqToken:      "",
	}
	err := snapshotEntity.GenerateSnapshotToken()
	if err != nil {
		return nil, err
	}
	return snapshotEntity, nil
}

//
//// 创建快照
//func (impl *ApplicationServiceImpl) CreateSnapshot(ctx context.Context, snapshotEntity *entity.TaskDtoSubmitSnapshot) error {
//	isLocked, err := repository.SnapshotRepo.CheckSnapshotTokenProcessing(ctx, snapshotEntity.ReqToken)
//	if isLocked {
//		return fmt.Errorf("snapshotToken processing!")
//	}
//	if err != nil {
//		return err
//	}
//	_ = repository.SnapshotRepo.SetSnapshotToken(ctx, snapshotEntity.ReqToken)
//	err = repository.SnapshotRepo.CreateSnapshot(ctx, snapshotEntity)
//	return err
//}

// 定义调用任务支持子域进行校验
func (impl *ApplicationServiceImpl) CheckSnapshotTask(ctx context.Context, snapshotEntity *entity.TaskDtoSubmitSnapshot) error {
	return nil
}

func (impl *ApplicationServiceImpl) CreateTask(ctx *context.Context, snapshotId int64) error {
	return nil
}
