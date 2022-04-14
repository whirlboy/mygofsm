package snapshot

import (
	"context"
	"mygofsm/my_order_project/application/service/snapshot_service"
	"mygofsm/my_order_project/domain/snapshot_domain/snapshot_aggregate/repository"
	"mygofsm/my_order_project/idl/gen-go/snapshot"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

func PreSubmitOrders(ctx context.Context, req snapshot.PreSubmitOrdersReq) (*snapshot.PreSubmitOrdersResp, error) {
	// 生成快照实体
	snapshotEntity, err := snapshot_service.GenerateSnapshotEntity(ctx, req)
	if err != nil {
		return nil, err
	}
	// 如果运行出错，删除token
	defer func() {
		if err != nil {
			repository.SnapshotRepo.DeleteSnapshotToken(ctx, snapshotEntity.ReqToken)
		}
	}()
	err = snapshot_service.CreateSnapshot(ctx, snapshotEntity)
	if err != nil {
		return nil, err
	}
	err = snapshot_service.CheckSnapshotTask(ctx, snapshotEntity)
	return nil, nil
}
