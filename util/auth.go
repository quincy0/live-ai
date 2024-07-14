package util

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/service/userService"
	"github.com/quincy0/qpro/qConfig"
)

const (
	UserIdKey   = "userid"
	UsernameKey = "username"
)

type User struct {
	UserId   int64
	Username string
}

var JWTAuth *jwt.GinJWTMiddleware

func init() {
	middleware, err := jwt.New(initParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	JWTAuth = middleware
}

func initParams() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(qConfig.Settings.Jwt.Secret),
		Timeout:         7 * 24 * time.Hour,
		MaxRefresh:      7 * 24 * time.Hour,
		PayloadFunc:     payloadFunc(),
		IdentityHandler: identityHandler(),
		Authenticator:   authenticator(),
		Authorizator:    authorizator(),
		Unauthorized:    unauthorized(),
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TimeFunc:        time.Now,
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*User); ok {
			return jwt.MapClaims{
				UserIdKey:   v.UserId,
				UsernameKey: v.Username,
			}
		}
		return jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		fmt.Println(claims)
		return &User{
			UserId:   int64(claims[UserIdKey].(float64)),
			Username: claims[UsernameKey].(string),
		}
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var params dto.LoginParam
		if err := c.ShouldBind(&params); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		username := params.Username
		password := params.Password

		if (username == "admin" && password == "admin") || (username == "test" && password == "test") {
			return &User{
				UserId:   100000,
				Username: username,
			}, nil
		}
		if userId := userService.Verify(c.Request.Context(), username, password); userId > 0 {
			return &User{
				UserId:   userId,
				Username: username,
			}, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}

// Authorizator
func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if _, ok := data.(*User); ok {
			return true
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

func ParseUser(c *gin.Context) *User {
	//claims := jwt.ExtractClaims(c)
	user, _ := c.Get(jwt.IdentityKey)
	return user.(*User)
}
