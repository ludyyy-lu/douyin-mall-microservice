package service_test

import (
	"context"
	"testing"

	service "github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/service"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/mocks"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestAddItemService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建 Mocks
	mockProductClient := mocks.NewMockClient(ctrl)
	mockCartStore := mocks.NewMockCartStore(ctrl)

	// 设置 Mock 预期
	mockProductClient.EXPECT().
		GetProduct(gomock.Any(), &product.GetProductReq{Id: 100}).
		Return(&product.GetProductResp{Product: &product.Product{Id: 100}}, nil)

	mockCartStore.EXPECT().
		AddItem(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	// 初始化服务
	svc := service.NewAddItemService(
		context.Background(),
		mockCartStore,
		mockProductClient, // 注入 Mock
	)

	// 执行测试
	resp, err := svc.Run(&cart.AddItemReq{
		UserId: 2001,
		Item:   &cart.CartItem{ProductId: 100, Quantity: 2},
	})

	// 验证结果
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestAddItemService_ProductNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductClient := mocks.NewMockClient(ctrl)
	mockProductClient.EXPECT().
		GetProduct(gomock.Any(), gomock.Any()).
		Return(&product.GetProductResp{Product: nil}, nil) // 返回空产品

	svc := service.NewAddItemService(
		context.Background(),
		nil, // CartStore 未被调用，可为 nil
		mockProductClient,
	)

	resp, err := svc.Run(&cart.AddItemReq{Item: &cart.CartItem{ProductId: 404}})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestAddItemService_RPCFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductClient := mocks.NewMockClient(ctrl)
	mockProductClient.EXPECT().
		GetProduct(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(500, "RPC error")) // 模拟 RPC 错误

	svc := service.NewAddItemService(
		context.Background(),
		nil,
		mockProductClient,
	)

	resp, err := svc.Run(&cart.AddItemReq{Item: &cart.CartItem{ProductId: 100}})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestAddItemService_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductClient := mocks.NewMockClient(ctrl)
	mockCartStore := mocks.NewMockCartStore(ctrl)

	mockProductClient.EXPECT().
		GetProduct(gomock.Any(), gomock.Any()).
		Return(&product.GetProductResp{Product: &product.Product{Id: 100}}, nil)

	mockCartStore.EXPECT().
		AddItem(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(gorm.ErrInvalidData) // 模拟数据库错误

	svc := service.NewAddItemService(
		context.Background(),
		mockCartStore,
		mockProductClient,
	)

	resp, err := svc.Run(&cart.AddItemReq{Item: &cart.CartItem{ProductId: 100}})
	assert.Error(t, err)
	assert.Nil(t, resp)
}
