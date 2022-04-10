package snapshot_service

import (
	"context"
	"fmt"
	"go-order-system/domain/snapshot_domain/snapshot_aggregate/entity"
	"go-order-system/domain/snapshot_domain/snapshot_aggregate/repository"
	"go-order-system/idl/gen-go/snapshot"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

func GenerateSnapshotEntity(ctx context.Context, req snapshot.PreSubmitOrdersReq) (*entity.TaskDtoSubmitSnapshot, error) {
	snapshotEntity := &entity.TaskDtoSubmitSnapshot{
		Id:            0,
		UserId:        0,
		SubmitType:    0,
		SubmitCommand: 0,
		ReqToken:      "",
	}
	err := snapshotEntity.GenerateSnapshotToken()
	if err != nil {
		return nil, err
	}
	return snapshotEntity, nil
}

// 创建快照
func CreateSnapshot(ctx context.Context, snapshotEntity *entity.TaskDtoSubmitSnapshot) error {
	isLocked, err := repository.SnapshotRepo.CheckSnapshotTokenProcessing(ctx, snapshotEntity.ReqToken)
	if isLocked {
		return fmt.Errorf("snapshotToken processing!")
	}
	if err != nil {
		return err
	}
	_ = repository.SnapshotRepo.SetSnapshotToken(ctx, snapshotEntity.ReqToken)
	err = repository.SnapshotRepo.CreateSnapshot(ctx, snapshotEntity)
	return err
}

// 定义调用任务支持子域进行校验
func CheckSnapshotTask(ctx context.Context, snapshotEntity *entity.TaskDtoSubmitSnapshot) error {
	return nil
}
