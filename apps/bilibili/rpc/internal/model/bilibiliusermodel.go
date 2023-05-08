package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BilibiliUserModel = (*customBilibiliUserModel)(nil)

type (
	// BilibiliUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBilibiliUserModel.
	BilibiliUserModel interface {
		bilibiliUserModel
		FindByDedeuserid(dedeuserid string) (*BilibiliUser, error)
		List(page, size int64) ([]*BilibiliUser, int64, error)
	}

	customBilibiliUserModel struct {
		*defaultBilibiliUserModel
	}
)

// NewBilibiliUserModel returns a model for the database table.
func NewBilibiliUserModel(conn sqlx.SqlConn) BilibiliUserModel {
	return &customBilibiliUserModel{
		defaultBilibiliUserModel: newBilibiliUserModel(conn),
	}
}

func (c *customBilibiliUserModel) FindByDedeuserid(dedeuserid string) (*BilibiliUser, error) {
	var res BilibiliUser
	err := c.conn.QueryRow(&res, fmt.Sprintf("select %s from %s where dedeuserid = ?", bilibiliUserRows, c.table), dedeuserid)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *customBilibiliUserModel) List(page, size int64) ([]*BilibiliUser, int64, error) {
	var res []*BilibiliUser
	err := c.conn.QueryRows(&res, fmt.Sprintf("select %s from %s order by  is_login desc, current_exp desc limit ?,?", bilibiliUserRows, c.table), (page-1)*size, size)
	if err != nil {
		return nil, 0, err
	}
	var count int64
	// 查询总页数
	err = c.conn.QueryRow(&count, fmt.Sprintf("select count(*) from %s", c.table))
	if err != nil {
		return nil, 0, err
	}
	total := count/size + 1
	return res, total, nil
}
