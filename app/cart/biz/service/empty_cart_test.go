package service_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/mocks"

	service "github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/service"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestEmptyCartService_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartStore := mocks.NewMockCartStore(ctrl)
	service := &service.EmptyCartService{
		CartStore: mockCartStore, // 假设已注入CartStore接口
		Ctx:       context.Background(),
	}

	t.Run("成功清空购物车", func(t *testing.T) {
		mockCartStore.EXPECT().
			EmptyCart(gomock.Any(), mysql.DB, 1001).
			Return(nil)

		resp, err := service.Run(&cart.EmptyCartReq{UserId: 1001})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("数据库操作失败", func(t *testing.T) {
		mockCartStore.EXPECT().
			EmptyCart(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(sql.ErrTxDone) // 模拟事务错误

		_, err := service.Run(&cart.EmptyCartReq{UserId: 1001})

		assert.Error(t, err)
		bizErr, ok := err.(*kerrors.BizStatusError)
		if ok {
			// 假设 kerrors.BizStatusError 有一个公开的 BizStatusCode 方法来获取错误码
			assert.Equal(t, 50001, bizErr.BizStatusCode())
		} else {
			t.Fatalf("expected error type *kerrors.BizStatusError, got %T", err)
		}
	})
}
