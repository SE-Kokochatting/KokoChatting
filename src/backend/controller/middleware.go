package controller

import (
	"KokoChatting/global"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Middleware struct{}

func (m *Middleware) ZapLogger() gin.HandlerFunc {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	path,err := global.GetGlobalConfig().GetConfigByPath("logger.ginlog")
	if err != nil{
		panic("config logger.ginlog get error: "+err.Error())
	}

	file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)

	fileWriteSyncer := zapcore.AddSync(file)
	core := zapcore.NewTee(
		//zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriteSyncer, global.GetLoggerLevel()),
	)
	logger := zap.New(core)
	return func(c *gin.Context){
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		end = end.UTC()

		fields := []zapcore.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.Error(e, fields...)
			}
		} else {
			logger.Info(path, fields...)
		}
	}
}

// JwtAuthValidate jwt身份信息验证
func (m *Middleware) JwtAuthValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 由于token保存在请求头中，所以需要使用c.getHeader，而不是c.Param函数
		tokenString := c.GetHeader("Authorization")
		// tokenString := c.Param("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			hmacSampleSecret := []byte(global.GetGlobalConfig().GetConfigByName("jwt.secret").(string))
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})
		if err != nil {
			c.JSON(404, gin.H{
				"status": 1001,
				"err": "unexpected signing method",
			})
			return
		}

		claims,ok := token.Claims.(jwt.MapClaims)
		//fmt.Println(claims)
		//claims, ok := token.Claims.(utilstruct.Claims)
		if  ok && token.Valid {
			c.Set("userUid", claims["Uid"])
			c.Set("userPassword", claims["Password"])
		}
		c.Next()
	}
}


func (m *Middleware) CORS()gin.HandlerFunc{
	return func(c *gin.Context){
		method := c.Request.Method //请求方法
		//fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")// 指明哪些请求源被允许访问资源，值可以为 "*"，"null"，或者单个源地址。
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")//对于预请求来说，指明了哪些头信息可以用于实际的请求中。
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")//对于预请求来说，哪些请求方式可以用于实际的请求。
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")//对于预请求来说，指明哪些头信息可以安全的暴露给 CORS API 规范的 API
		c.Header("Access-Control-Allow-Credentials", "true")//指明当请求中省略 creadentials 标识时响应是否暴露。对于预请求来说，它表明实际的请求中可以包含用户凭证。
		
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}