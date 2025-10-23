package mw

import (
	"admin/rpc/client"
	"common/consts"
	"common/pkg/errno"
	"common/pkg/utils"
	"context"
	"gen/kitex_gen/base"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
	"github.com/pkg/errors"
	"time"
)

type jwtLogin struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Captcha  string `form:"captcha" json:"captcha"`
}

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJwt(enforcer *casbin.Enforcer) {
	hlog.Info("Jwt")
	JwtMiddleware, _ = jwt.New(
		&jwt.HertzJWTMiddleware{
			Realm:       "deer",
			Key:         []byte(consts.SecretKey),
			Timeout:     time.Hour,
			MaxRefresh:  time.Hour,
			TimeFunc:    time.Now,
			IdentityKey: consts.IdentityKey,
			IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
				claims := jwt.ExtractClaims(ctx, c)
				payloadMap, _ := claims[consts.IdentityKey].(map[string]interface{})
				return payloadMap
			},

			Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
				var err error
				var req jwtLogin

				if err = c.BindAndValidate(&req); err != nil {
					return "", err
				}
				if len(req.Username) == 0 || len(req.Password) == 0 {
					return "", jwt.ErrMissingLoginValues
				}
				resp, err := client.UserClient.LoginUser(ctx, &base.CheckAccountReq{
					Username: req.Username,
					Password: req.Password,
					Captcha:  req.Captcha,
				})

				if err != nil {
					return nil, err
				}
				return map[string]interface{}{
					"id": resp.Data.GetId(),
				}, nil
			},
			Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
				obj := string(c.URI().Path())
				act := string(c.Method())
				hlog.Infof("obj: %v | act: %v", obj, act)
				// IdentityHandler 返回的数据
				if v, ok := data.(map[string]interface{}); ok {
					hlog.Infof("authorizator data: %v", v)
					return true
				}

				//existToken := rpc.NewToken(ctx, c).IsExistByUserId(int64(id))
				//if !existToken {
				//	return false
				//}
				// check the role status
				//roleInfo, err := rpc.NewRole(ctx, c).RoleInfoByID(cast.ToInt64(roleId))
				//// if the role is not exist or the role is not active, return false
				//if err != nil {
				//	hlog.Error(err, "role is not exist")
				//	return false
				//}

				//if roleInfo.Status != 1 {
				//	hlog.Error("role cache is not a valid *ent.Role or the role is not active")
				//	return false
				//}

				//sub := roleId
				//check the permission
				//pass, err := enforcer.Enforce(sub, obj, act)
				//if err != nil {
				//	hlog.Error("casbin err,  role id: ", roleId, " path: ", obj, " method: ", act, " pass: ", pass, " err: ", err.Error())
				//	return false
				//}
				//if !pass {
				//	hlog.Info("casbin forbid role id: ", roleId, " path: ", obj, " method: ", act, " pass: ", pass)
				//}
				//hlog.Info("casbin allow role id: ", roleId, " path: ", obj, " method: ", act, " pass: ", pass)
				//return pass

				return true
			},
			PayloadFunc: func(data interface{}) jwt.MapClaims {

				if v, ok := data.(map[string]interface{}); ok {
					return jwt.MapClaims{
						consts.IdentityKey: v,
					}
				}
				return jwt.MapClaims{}
			},
			Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
				utils.SendResponse(c, errno.NewErrNo(10002, "您没有访问此资源的权限"), message, 0, "")
			},
			LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
				utils.SendResponse(c, errno.Success,
					map[string]interface{}{
						"token":  token,
						"expire": expire.Format(time.DateTime),
					}, 0, "")
			},
			RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
				utils.SendResponse(c, errno.Success,
					map[string]interface{}{
						"token":  token,
						"expire": expire.Format(time.DateTime),
					}, 0, "")
			},
			LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {

				var err error
				//err = func...
				if err != nil {
					utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
				}
				utils.SendResponse(c, errno.Success, nil, 0, "")

			},

			HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
				var expr errno.ErrNo
				switch {
				case errors.As(e, &expr):
					return expr.ErrMsg
				default:
					return expr.Error()
				}
			},
			ParseOptions: []jwtv4.ParserOption{jwtv4.WithJSONNumber()},
		},
	)

}
