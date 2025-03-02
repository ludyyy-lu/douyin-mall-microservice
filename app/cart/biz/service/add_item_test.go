package service_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"
	service "github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/service"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/mocks"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestAddItemService_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建mock对象
	mockProductClient := mocks.NewMockClient(ctrl)
	mockCartStore := mocks.NewMockCartStore(ctrl)

	// 初始化服务，注入mock依赖
	service := &service.AddItemService{
		Ctx:           context.Background(),
		ProductClient: mockProductClient,
		CartStore:     mockCartStore,
	}

	t.Run("成功添加商品", func(t *testing.T) {
		// 设置mock预期
		mockProductClient.EXPECT().GetProduct(gomock.Any(), &product.GetProductReq{Id: 123}).
			Return(&product.GetProductResp{
				Product: &product.Product{Id: 123},
			}, nil)

		mockCartStore.EXPECT().AddItem(gomock.Any(), mysql.DB, &model.Cart{
			UserID:    1001,
			ProductID: 123,
			Qty:       2,
		}).Return(nil)

		// 执行测试
		resp, err := service.Run(&cart.AddItemReq{
			UserId: 1001,
			Item:   &cart.CartItem{ProductId: 123, Quantity: 2},
		})

		// 验证结果
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("商品不存在", func(t *testing.T) {
		mockProductClient.EXPECT().GetProduct(gomock.Any(), gomock.Any()).
			Return(&product.GetProductResp{Product: nil}, nil)

		_, err := service.Run(&cart.AddItemReq{Item: &cart.CartItem{ProductId: 456}})
		assert.Error(t, err)
		//assert.Equal(t, 40004, err.(*kerrors.BizStatusError).code)
		bizErr, ok := err.(*kerrors.BizStatusError)
		if ok {
			// 假设 kerrors.BizStatusError 有一个公开的 BizStatusCode 方法来获取错误码
			assert.Equal(t, 40004, bizErr.BizStatusCode())
		} else {
			t.Fatalf("expected error type *kerrors.BizStatusError, got %T", err)
		}
	})

	t.Run("数据库操作失败", func(t *testing.T) {
		mockProductClient.EXPECT().GetProduct(gomock.Any(), gomock.Any()).
			Return(&product.GetProductResp{Product: &product.Product{Id: 123}}, nil)

		mockCartStore.EXPECT().AddItem(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(sql.ErrConnDone) // 模拟数据库错误

		_, err := service.Run(&cart.AddItemReq{Item: &cart.CartItem{ProductId: 123}})
		assert.Error(t, err)
		//assert.Equal(t, 50000, err.(*kerrors.BizStatusError).code)
		bizErr, ok := err.(*kerrors.BizStatusError)
		if ok {
			// 假设 kerrors.BizStatusError 有一个公开的 BizStatusCode 方法来获取错误码
			assert.Equal(t, 50000, bizErr.BizStatusCode())
		} else {
			t.Fatalf("expected error type *kerrors.BizStatusError, got %T", err)
		}
	})
}
