package utils

import (
	"common/consts"
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
)

func GetTokenId(ctx context.Context, c *app.RequestContext) int64 {
	claims := jwt.ExtractClaims(ctx, c)
	if len(claims) == 0 {
		return 0
	}
	token, _ := claims[consts.IdentityKey].(map[string]interface{})
	hlog.Infof("token id: %d", token)
	id, err := token["id"].(json.Number).Int64()
	if err != nil {
		return 0
	}
	return id
}
