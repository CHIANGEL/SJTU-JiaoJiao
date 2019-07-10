package mock

import (
	"context"
	"errors"
	user "jiaojiao/srv/user/proto"

	"github.com/micro/go-micro/client"
)

type mockUserSrv struct{}
type mockAdminSrv struct{}

func (a *mockUserSrv) Create(ctx context.Context, req *user.UserCreateRequest, opts ...client.CallOption) (*user.UserCreateResponse, error) {
	var rsp user.UserCreateResponse
	if req.StudentId == "" || req.StudentName == "" {
		rsp.Status = user.UserCreateResponse_INVALID_PARAM
	} else {
		if req.StudentId == "1000" {
			rsp.Status = user.UserCreateResponse_SUCCESS
			rsp.UserId = 1
		} else if req.StudentId == "2000" {
			return &rsp, errors.New("")
		} else {
			rsp.Status = user.UserCreateResponse_USER_EXIST
			rsp.UserId = 1
		}
	}
	return &rsp, nil
}

func (a *mockUserSrv) Query(ctx context.Context, req *user.UserQueryRequest, opts ...client.CallOption) (*user.UserInfo, error) {
	var rsp user.UserInfo
	if req.UserId != 0 {
		if req.UserId == 1000 {
			rsp.UserId = 1000
			rsp.UserName = "test"
			rsp.AvatarId = "5d23ea2c32311335f935cd14"
			rsp.Telephone = "12345678901"
			rsp.StudentId = "1000"
			rsp.StudentName = "jiang"
		} else if req.UserId == 2000 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

func (a *mockUserSrv) Find(ctx context.Context, req *user.UserFindRequest, opts ...client.CallOption) (*user.UserFindResponse, error) {
	var rsp user.UserFindResponse
	return &rsp, nil
}

func (a *mockAdminSrv) Create(ctx context.Context, req *user.AdminUserRequest, opts ...client.CallOption) (*user.AdminUserResponse, error) {
	var rsp user.AdminUserResponse
	if req.StudentId == "" {
		rsp.Status = user.AdminUserResponse_INVALID_PARAM
	} else {
		if req.StudentId == "1001" {
			rsp.Status = user.AdminUserResponse_SUCCESS
			rsp.AdminId = 1
		} else if req.StudentId == "2001" {
			return &rsp, errors.New("")
		} else {
			rsp.Status = user.AdminUserResponse_USER_EXIST
			rsp.AdminId = 1
		}
	}
	return &rsp, nil
}

func (a *mockAdminSrv) Find(ctx context.Context, req *user.AdminUserRequest, opts ...client.CallOption) (*user.AdminUserResponse, error) {
	var rsp user.AdminUserResponse
	if req.StudentId == "" {
		rsp.Status = user.AdminUserResponse_INVALID_PARAM
	} else {
		if req.StudentId == "1001" {
			rsp.Status = user.AdminUserResponse_SUCCESS
			rsp.AdminId = 1
		} else if req.StudentId == "2001" {
			return &rsp, errors.New("")
		} else {
			rsp.Status = user.AdminUserResponse_NOT_FOUND
		}
	}
	return &rsp, nil
}

func NewAdminUserService() user.AdminUserService {
	return new(mockAdminSrv)
}

func NewUserService() user.UserService {
	return new(mockUserSrv)
}