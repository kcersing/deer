package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"member/biz/dal/db/ent"
)

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		hlog.Error("警告！回滚失败:", rerr)
		err = fmt.Errorf("%w: %v", err, rerr)
		hlog.Error("失败原因:", err, rerr)
	}
	return err
}
