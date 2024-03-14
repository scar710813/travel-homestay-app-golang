package svc

import (
	"homestay/app/travel/cmd/rpc/internal/config"
	"homestay/app/travel/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	HomestayModel        model.HomestayModel
	UserHomestayModel    model.UserHomestayModel
	HistoryModel         model.HistoryModel
	HistoryHomestayModel model.HistoryHomestayModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		// 我了个骚刚bug: 没有初始化HomestayModel导致，travel模块的RPC调不动model
		HomestayModel:        model.NewHomestayModel(sqlConn, c.Cache),
		UserHomestayModel:    model.NewUserHomestayModel(sqlConn, c.Cache),
		HistoryModel:         model.NewHistoryModel(sqlConn, c.Cache),
		HistoryHomestayModel: model.NewHistoryHomestayModel(sqlConn, c.Cache),
	}
}
