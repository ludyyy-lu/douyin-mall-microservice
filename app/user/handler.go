// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/service"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserServiceImpl 实现了两个方法：Register 和 Login

// Register implements the UserServiceImpl interface.
// Register 处理用户注册请求。
// 它接收一个上下文和一个注册请求，并返回注册响应和可能的错误。
// `(s *UserServiceImpl)`：这是方法的接收者，表示该方法属于 `UserServiceImpl` 结构体的指针类型。
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 创建一个新的注册服务实例，并执行注册逻辑
	resp, err = service.NewRegisterService(ctx).Run(req)

	// 返回注册响应和错误
	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}
