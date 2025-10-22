package mw

import (
	"admin/rpc/client"
	"common/consts"
	"common/pkg/errno"
	"common/pkg/utils"
	"context"
	"encoding/json"
	"gen/kitex_gen/base"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
	"github.com/pkg/errors"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

type jwtLogin struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Captcha  string `form:"captcha" json:"captcha"`
}

func InitJwt() {
	JwtMiddleware, _ = jwt.New(
		&jwt.HertzJWTMiddleware{
			Key:        []byte(consts.SecretKey),
			Timeout:    time.Hour,
			MaxRefresh: time.Hour,
			Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
				var err error
				var req jwtLogin

				if err = c.Bind(&req); err != nil {
					hlog.Info(err)
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
				hlog.Info(resp)
				payLoadMap := make(map[string]interface{})
				//payLoadMap[consts.IdentityKey] = strconv.Itoa(int(resp.Data.GetId()))
				//payLoadMap["roleIds"] = resp.Data.GetRoles()
				return payLoadMap, nil
			},
			Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {

				obj := string(c.URI().Path())
				act := string(c.Method())
				hlog.Info(obj, act)
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
				if v, ok := data.(int64); ok {
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
						"expire": expire.Format(time.RFC3339),
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
			IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
				claims := jwt.ExtractClaims(ctx, c)
				userid, _ := claims[consts.IdentityKey].(json.Number).Int64()
				return userid
			},
			IdentityKey:   consts.IdentityKey,
			TokenLookup:   "header: Authorization, query: token, cookie: jwt",
			TokenHeadName: "Bearer",

			TimeFunc: time.Now,
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
