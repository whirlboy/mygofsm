package repository

import (
	"context"
	"go-order-system/domain/snapshot_domain/snapshot_aggregate/entity"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

var SnapshotRepo SnapshotRepository

type SnapshotRepository interface {
	GetSnapshotById(ctx context.Context, snapshotId int64) (*entity.TaskDtoSubmitSnapshot, error)
	CheckSnapshotTokenProcessing(ctx context.Context, snapshotToken string) (isLocked bool, err error)
	SetSnapshotToken(ctx context.Context, snapshotToken string) error
	DeleteSnapshotToken(ctx context.Context, snapshotToken string) error
	CreateSnapshot(ctx context.Context, snapshotEntity *entity.TaskDtoSubmitSnapshot) error
}
