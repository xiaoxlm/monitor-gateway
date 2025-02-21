package middleware

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/lie-flat-planet/httputil"
	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/pkg/authorization"
	"net/http"
	"strings"
)

func BasicAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httputil.ErrorRESP{
				Code: http.StatusUnauthorized*1e6 + config.Config.Server.Code,
				Msg:  "Authorization header required",
			})
			return
		}

		token := authorization.ParseAuthorization(authHeader)

		username, password, ok := decodeBasicAuth(token.Get("Basic"))
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httputil.ErrorRESP{
				Code: http.StatusUnauthorized*1e6 + config.Config.Server.Code,
				Msg:  `token is invalid`,
			})
			return
		}

		if username == config.Config.ClientID && password == config.Config.ClientSecret {
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, httputil.ErrorRESP{
			Code: http.StatusUnauthorized*1e6 + config.Config.Server.Code,
			Msg:  "check token failed",
		})
		return
	}
}

func decodeBasicAuth(encodedStr string) (username, password string, ok bool) {
	c, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
