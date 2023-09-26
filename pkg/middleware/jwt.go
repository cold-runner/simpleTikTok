package middleware

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/simpleTikTok/apiServer/biz/model/ApiServer"
	"github.com/cold-runner/simpleTikTok/apiServer/rpc"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware
var identityKey = "id"
var IdentityKey = identityKey

func InitJwt() {
	log.Infow("Initializing Jwt middleware")

	// the jwt middleware
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		SigningAlgorithm: "HS256",
		Key:              []byte("secret key"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,

		// 登录时触发的方法，用于认证用户登录
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var req ApiServer.UserLoginRequest
			var id int64
			var err error
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			// 调用 rpc 方法进行登录
			id, err = rpc.Login(context.Background(),
				&UserService.UserLoginRequest{
					Username: req.Username,
					Password: req.Password,
				})
			c.Set(identityKey, id)
			return id, err
		},
		//Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
		//	if v, ok := data.(*User); ok && v.UserName == "admin" {
		//		return true
		//	}
		//
		//	return false
		//},

		// 添加了自定义负载信息的函数，这里是将用户id存入 token 的 payload 部分
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// 登录失败时的响应
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			err := errno.ErrUnauthorized
			c.JSON(code, map[string]interface{}{
				"status_code": err.HTTP,
				"status_msg":  err.Message,
				"user_id":     -1,
				"token":       "",
			})
		},
		// 登录成功时的响应
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			v, _ := c.Get("id")
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": 0,
				"status_msg":  "Login successfully.",
				"user_id":     v,
				"token":       token,
			})
		},
		//LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
		//	c.JSON(http.StatusOK, map[string]interface{}{
		//		"code": http.StatusOK,
		//	})
		//},
		//RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
		//	c.JSON(http.StatusOK, map[string]interface{}{
		//		"code":   http.StatusOK,
		//		"token":  token,
		//		"expire": expire.Format(time.RFC3339),
		//	})
		//},
		// 在每次登录成功后的请求中，从 token 中提取用户信息
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &ApiServer.User{
				Id: int64(claims[identityKey].(float64)),
			}
		},
		IdentityKey: identityKey,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, form: token, param: token, cookie: token",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName: "Bearer",
		//WithoutDefaultTokenHeadName: false,
		TimeFunc: time.Now,
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			return e.Error()
		},
		//SendCookie:        true,
		//CookieMaxAge:      time.Hour,
		//SecureCookie:      false,
		//CookieHTTPOnly:    false,
		//CookieDomain:      ".test.com",
		//CookieName:        "jwt-cookie",
		//CookieSameSite:    protocol.CookieSameSiteDisabled,
		//SendAuthorization: true,
		//DisabledAbort:     false,
	})
}

func MiddlewareFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		middlewareFunc(ctx, c) // 调用实现JWT认证逻辑的函数
	}
}

func middlewareFunc(ctx context.Context, c *app.RequestContext) {
	claims, err := JwtMiddleware.GetClaimsFromJWT(ctx, c)
	log.Debugw("get jwt claims", "claims", claims)
	if err != nil {
		if err == jwt.ErrEmptyQueryToken || err == jwt.ErrEmptyCookieToken {
			log.Errorw("feed request with empty token")
			c.Next(ctx)
			return
		}
		log.Errorw("get jwt claims failed", "err", err)
		unauthorized(ctx, c, http.StatusUnauthorized, err)
		return
	}

	exp, exists := claims["exp"]
	if !exists {
		log.Errorw("missing exp field")
		unauthorized(ctx, c, http.StatusBadRequest, jwt.ErrMissingExpField)
		return
	}

	if err = checkTokenExpiration(exp); err != nil {
		log.Errorw("check token expiration failed", "err", err)
		unauthorized(ctx, c, http.StatusUnauthorized, err)
		return
	}

	log.Debugw("set jwt payload", "claims", claims)
	c.Set("JWT_PAYLOAD", claims)
	identity := JwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(JwtMiddleware.IdentityKey, identity)
	}

	if !JwtMiddleware.Authorizator(identity, ctx, c) {
		log.Errorw("jwt authorizator failed", "err", jwt.ErrForbidden)
		unauthorized(ctx, c, http.StatusForbidden, jwt.ErrForbidden)
		return
	}

	c.Next(ctx)

}

func checkTokenExpiration(exp interface{}) error {
	var expInt64 int64
	switch v := exp.(type) {
	case float64:
		expInt64 = int64(v)
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return jwt.ErrWrongFormatOfExp
		}
		expInt64 = n
	default:
		return jwt.ErrWrongFormatOfExp
	}
	if expInt64 < JwtMiddleware.TimeFunc().Unix() {
		return jwt.ErrExpiredToken
	}
	return nil
}

func unauthorized(ctx context.Context, c *app.RequestContext, code int, err error) {
	message := JwtMiddleware.HTTPStatusMessageFunc(err, ctx, c)
	c.Header("WWW-Authenticate", "JWT realm="+JwtMiddleware.Realm)
	if !JwtMiddleware.DisabledAbort {
		c.Abort()
	}
	JwtMiddleware.Unauthorized(ctx, c, code, message)
}
