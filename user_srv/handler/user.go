/**
@auther:liuzhijun
@data:2023/4/3
**/
package handler

import (
	"context"
	"gorm.io/gorm"
	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/model"
	user "mxshop_srvs/user_srv/proto/gen"
)

type UserServer struct{}

func ModelToResponse(u model.User) user.UserInfoResponse {
	//在grpc的message中字段有默认值，你不能随便赋值nil,容易出错
	//得搞清楚，哪些字段有默认值
	userInfoRsp := user.UserInfoResponse{
		Id:       u.ID,
		PassWord: u.Password,
		NickName: u.NickName,
		Gender:   u.Gender,
		Role:     int32((u.Role)),
	}
	if u.Birthday != nil {
		userInfoRsp.BirthDay = uint64(u.Birthday.Unix())
	}
	return userInfoRsp
}
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserServer) GetUserList(ctx context.Context, req *user.PageInfo) (*user.UserListResponse, error) {
	//获取用户列表
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &user.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)
	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, u := range users {
		userInfoRsp := ModelToResponse(u)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}
