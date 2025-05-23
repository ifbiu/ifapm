package dao

import (
	"context"
	"database/sql"
	"ifapm"
)

type skuDao struct{}

var SkuDao = &skuDao{}

func (s *skuDao) Get(ctx context.Context, id int64) map[string]interface{} {
	info := ifapm.DBUtil.QueryFirst(ifapm.Infra.Db.QueryContext(ctx, "SELECT * FROM t_sku WHERE id=?", id))
	if len(info) == 0 {
		return nil
	}
	return info
}

func (s *skuDao) Decr(ctx context.Context, id int64, num int32) (sql.Result, error) {
	return ifapm.Infra.Db.ExecContext(ctx, "update t_sku set num = num - ? where id=? and (num - ?) >= 0", num, id)
}
