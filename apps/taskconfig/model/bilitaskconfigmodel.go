package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BiliTaskConfigModel = (*customBiliTaskConfigModel)(nil)

type (
	// BiliTaskConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBiliTaskConfigModel.
	BiliTaskConfigModel interface {
		biliTaskConfigModel
		GetConfigByDedeuserid(dedeuserid string) (*BiliTaskConfig, error)
	}

	customBiliTaskConfigModel struct {
		*defaultBiliTaskConfigModel
	}
)

// NewBiliTaskConfigModel returns a model for the database table.
func NewBiliTaskConfigModel(conn sqlx.SqlConn) BiliTaskConfigModel {
	return &customBiliTaskConfigModel{
		defaultBiliTaskConfigModel: newBiliTaskConfigModel(conn),
	}
}

func (c *customBiliTaskConfigModel) GetConfigByDedeuserid(dedeuserid string) (*BiliTaskConfig, error) {
	var res BiliTaskConfig
	err := c.conn.QueryRow(&res, fmt.Sprintf("select %s from %s where dedeuserid = ?", biliTaskConfigRows, c.table), dedeuserid)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
