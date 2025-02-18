package main

import (
	"context"
	"errors"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/global"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/repository"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/repository/model"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if req.Address == nil {
		logrus.Errorln("address is empty")
		return nil, errors.New("address is empty")
	}
	orderID := uuid.NewString()
	orderData := model.Order{
		OrderID:      orderID,
		UserID:       req.UserId,
		UserCurrency: req.UserCurrency,
		Address: model.Address{
			Email:         req.Email,
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
	}
	if len(req.OrderItems) == 0 {
		logrus.Errorln("order items is empty")
		return nil, errors.New("order items is empty")
	}
	for _, item := range req.OrderItems {
		orderData.OrderItems = append(orderData.OrderItems, model.OrderItem{
			ProductID: item.Item.ProductId,
			Quantity:  item.Item.Quantity,
			Cost:      item.Cost,
		})
	}
	if err := global.DB.Create(&orderData).Error; err != nil {
		logrus.Errorln("create order failed")
		return nil, errors.New("create order failed")
	}
	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: orderID,
		},
	}
	return resp, nil
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	userID := req.GetUserId()
	orderDB := repository.NewOrderRepo(global.DB)
	orders, err := orderDB.ListOrders(userID)
	if err != nil {
		logrus.Errorln("list order failed")
		return nil, errors.New("list order failed")
	}
	logrus.Debugln(orders)
	resp = &order.ListOrderResp{
		Orders: make([]*order.Order, 0),
	}
	for _, orderData := range orders {
		orderItems := make([]*order.OrderItem, 0)
		for _, item := range orderData.OrderItems {
			orderItems = append(orderItems, &order.OrderItem{
				Cost: item.Cost,
				Item: &cart.CartItem{
					ProductId: item.ProductID,
					Quantity:  item.Quantity,
				},
			})
		}
		resp.Orders = append(resp.Orders, &order.Order{
			OrderId:      orderData.OrderID,
			UserId:       orderData.UserID,
			UserCurrency: orderData.UserCurrency,
			Email:        orderData.Email,
			Address: &order.Address{
				StreetAddress: orderData.Address.StreetAddress,
				City:          orderData.Address.City,
				State:         orderData.Address.State,
				Country:       orderData.Address.Country,
				ZipCode:       orderData.Address.ZipCode,
			},
			OrderItems: orderItems,
		})

	}

	return resp, nil
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// TODO: Your code here...
	return
}
