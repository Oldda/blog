package middlewares

import(
	"github.com/gin-gonic/gin"
	"blog/util"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("Authorization")
		token = string([]byte(token)[len("Bearer "):])
		if token == "" {
			c.JSON(util.NewApiResp("AUTHORIZATION_TOKEN_LOST"))
			c.Abort()
			return
		}
		j := util.NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		claims, err := j.ParserToken(token)

		if err != "" {
			c.JSON(util.NewApiResp(err))
			c.Abort()
			return
		}

		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)

	}
}