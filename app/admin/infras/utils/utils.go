package utils

import (
	"common/consts"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func GetTokenId(c *app.RequestContext) int64 {
	id, exist := c.Get(consts.IdentityKey)
	if exist || id != nil {
		uId, ok := id.(string)
		if ok {
			uid, err := strconv.ParseInt(uId, 10, 64)
			if err != nil {
				return 0
			}
			return uid
		}
		return 0
	}
	return 0
}
