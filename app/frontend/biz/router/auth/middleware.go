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

// Code generated by hertz generator.

package auth

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _authMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _logoutMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

// 用户相关的中间件

func _userMw() []app.HandlerFunc {
	// 用户相关路由的中间件
	return []app.HandlerFunc{middleware.Auth()}
}

func _getuserprofileMw() []app.HandlerFunc {
	// 获取用户信息的中间件
	return nil
}

func _profileMw() []app.HandlerFunc {
	// 用户个人信息页面相关中间件
	return nil
}

func _updateuserprofileMw() []app.HandlerFunc {
	// 更新用户信息的中间件
	return nil
}

func _deleteuserMw() []app.HandlerFunc {
	// 删除用户的中间件
	return nil
}
