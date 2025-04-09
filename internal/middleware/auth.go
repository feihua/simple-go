package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/feihua/simple-go/pkg/config"
	"github.com/feihua/simple-go/pkg/redis"
	"github.com/feihua/simple-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JwtMiddleware JWT中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		utils.Logger.Debugf("请求url: %s", path)
		if path == "/api/system/user/login" {
			c.Next()
			return
		}
		// 获取认证头
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.GlobalAppConfig.Jwt.AccessSecret), nil
		})

		if err != nil || !tkn.Valid {
			utils.Logger.Debugf("Invalid: %s", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userId := claims["userId"].(float64)
		userName := claims["userName"].(string)
		c.Set("userId", userId)
		c.Set("userName", userName)

		ctx := context.Background()
		var key = fmt.Sprintf("simple:apiUrl:%d", int64(userId))
		result, err := redis.Rdb.Get(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no permissions"})
			c.Abort()
			return
		}
		apiUrls := strings.Split(result, ",")
		var flag = false
		path = strings.Split(path, "?")[0]
		for _, url := range apiUrls {
			if path == url {
				flag = true
				break
			}
		}

		if flag {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no permissions"})
			c.Abort()
			return
		}
	}
}
