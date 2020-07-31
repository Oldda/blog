package middlewares

import(
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"blog/util"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(302, gin.H{
				"status": false,
				"response_code": 1001,
				"msg":    "请求未携带token，无权限访问",
				"data":   nil,
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		// 初始化一个JWT对象实例，并根据结构体方法来解析token
		j := util.NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		claims, err := j.ParserToken(token)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"response_code" : 1003,
				"msg":    err.Error(),
				"data":   nil,
			})
			c.Abort()
			return
		}

		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)

	}
}