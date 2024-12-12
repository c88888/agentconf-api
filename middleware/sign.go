package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

func SignAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sign := c.GetHeader("X-Sign")
		ts := c.Query("timestamp")
		agentHostname := c.Query("agent_hostname")
		agentHeader := c.GetHeader("X-agent")
		if sign == "" || ts == "" || agentHostname == "" || agentHeader != "categraf" {
			response.NoAuth("非法参数", c)
			c.Abort()
			return
		}

		tsInt, err := strconv.ParseInt(ts, 10, 64)
		if err != nil {
			response.NoAuth("非法参数ts", c)
			c.Abort()
			return
		}

		currentTime := time.Now().Unix()
		if currentTime-tsInt > 300 {
			response.NoAuth("请求已过期", c)
			c.Abort()
			return
		}
		signkey := global.GVA_CONFIG.BasicAuth.User + global.GVA_CONFIG.BasicAuth.Password
		toSign := fmt.Sprintf("%s-sign-key=%s@&target=%s@&sign-time=%s", ts, signkey, agentHostname, ts)
		// 生成 HMAC-SHA256 签名
		mac := hmac.New(sha256.New, []byte(signkey))
		mac.Write([]byte(toSign))
		calculatedSign := hex.EncodeToString(mac.Sum(nil))
		if sign != calculatedSign {
			response.NoAuth("非法访问", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
