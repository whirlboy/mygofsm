package snapshot_aggregate

import (
	"code.waizixi.com/go-order-system/domain/snapshot_domain/snapshot_aggregate/entity"
	"context"
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
