package infras

import (
	"context"
	"order/biz/dal/db"
	"order/biz/dal/db/ent"

	"github.com/pkg/errors"
)

// WithTx 是一个辅助函数，用于封装数据库事务操作。
// 它会自动处理事务的提交和回滚。
func WithTx(fn func(tx *ent.Tx) error) (err error) {

	tx, err := db.Client.Tx(context.Background())
	if err != nil {
		return errors.Wrap(err, "创建事务失败")
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			// 重新抛出 panic，以便上层可以捕获
			panic(r)
		} else if err != nil {
			// 如果发生错误，回滚事务
			if rbErr := tx.Rollback(); rbErr != nil {
				err = errors.Wrapf(err, "回滚事务失败: %v", rbErr)
			}
		} else {
			// 如果没有错误，提交事务
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
