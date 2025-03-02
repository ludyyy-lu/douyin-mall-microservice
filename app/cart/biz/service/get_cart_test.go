package service_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/service"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/mocks"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
	//"go.uber.org/mock/gomock"
)

func TestGetCartService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockCartStore(ctrl)
	Ctx := context.Background()
	mockDB := &gorm.DB{} // 假设的数据库连接，测试时可使用空对象或mock

	t.Run("正常获取购物车", func(t *testing.T) {
		// 准备模拟数据
		expectedItems := []*model.Cart{
			{ProductID: 1, Qty: 2},
			{ProductID: 2, Qty: 3},
		}

		// 设置Mock预期（关键修改点）
		mockStore.EXPECT().
			GetCartByUserId(Ctx, mockDB, int64(1001)). // 补全三个参数
			Return(expectedItems, nil)

		// 初始化服务（需要确保服务包含DB字段）
		svc := &service.GetCartService{
			CartStore: mockStore,
			DB:        mockDB, // 需要确认服务结构体是否有这个字段
			Ctx:       Ctx,
		}

		// 执行测试
		resp, err := svc.Run(&cart.GetCartReq{UserId: 1001})

		// 验证结果
		assert.NoError(t, err)
		assert.Len(t, resp.Items, 2)
		assert.Equal(t, int64(1), resp.Items[0].ProductId)
	})

	t.Run("数据库错误处理", func(t *testing.T) {
		// 使用gomock.Any()匹配参数
		mockStore.EXPECT().
			GetCartByUserId(
				gomock.Any(), // context
				gomock.Any(), // *gorm.DB
				int64(1002),  // 明确类型
			).
			Return(nil, sql.ErrConnDone)

		svc := &service.GetCartService{
			CartStore: mockStore,
			DB:        mockDB,
			Ctx:       Ctx,
		}

		_, err := svc.Run(&cart.GetCartReq{UserId: 1002})

		// 验证错误处理
		var bizErr *kerrors.BizStatusError
		if assert.ErrorAs(t, err, &bizErr) {
			assert.Equal(t, 50002, bizErr.BizStatusCode)
			assert.Contains(t, bizErr.Error(), "获取购物车失败")
		}
	})
}
