package snapshot_aggregate

import (
	"context"
	"mygofsm/my_order_project/domain/snapshot_domain/snapshot_aggregate/entity"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

type SnapshotRepositoryImpl struct {
}

func (impl *SnapshotRepositoryImpl) CheckSnapshotTokenProcessing(ctx context.Context, snapshotToken string) (isLocked bool, err error) {
	return true, nil
}

func (impl *SnapshotRepositoryImpl) GetSnapshotById(ctx context.Context, snapshotId int64) (*entity.TaskDtoSubmitSnapshot, error) {
	return nil, nil
}
func (impl *SnapshotRepositoryImpl) SetSnapshotToken(ctx context.Context, snapshotToken string) error {
	return nil
}
func (impl *SnapshotRepositoryImpl) DeleteSnapshotToken(ctx context.Context, snapshotToken string) error {
	return nil
}

func (impl *SnapshotRepositoryImpl) CreateSnapshot(ctx context.Context, snapshotEntity *entity.TaskDtoSubmitSnapshot) error {
	return nil
}
