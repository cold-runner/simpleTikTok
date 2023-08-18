package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func Init() {
	// the jwt middleware
	JwtMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            "simpleTikTok",
		SigningAlgorithm: "HS256",
		Key:              []byte("secret key"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals login
			if err := c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Hertz",
					FirstName: "CloudWeGo",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": http.StatusOK,
			})
		},
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &User{
				UserName: claims[identityKey].(string),
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
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName:               "Bearer",
		WithoutDefaultTokenHeadName: false,
		TimeFunc:                    time.Now,
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			return e.Error()
		},
		SendCookie:        true,
		CookieMaxAge:      time.Hour,
		SecureCookie:      false,
		CookieHTTPOnly:    false,
		CookieDomain:      ".test.com",
		CookieName:        "jwt-cookie",
		CookieSameSite:    protocol.CookieSameSiteDisabled,
		SendAuthorization: true,
		DisabledAbort:     false,
	})
}
