package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PushConfigModel = (*customPushConfigModel)(nil)

type (
	// PushConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPushConfigModel.
	PushConfigModel interface {
		pushConfigModel
	}

	customPushConfigModel struct {
		*defaultPushConfigModel
	}
)

// NewPushConfigModel returns a model for the database table.
func NewPushConfigModel(conn sqlx.SqlConn) PushConfigModel {
	return &customPushConfigModel{
		defaultPushConfigModel: newPushConfigModel(conn),
	}
}
