package snapshot

import (
	"context"
	"testing"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

func Test_Division_1(t *testing.T) {

	ctx := context.TODO()
	OrderServiceFacade.PreCreateOrder(&ctx, nil)
}

func Test_Division_2(t *testing.T) {
	t.Error("就是不通过")
}
