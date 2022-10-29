package controller

import (
	"KokoChatting/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Middleware struct{
	baseController
}

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
		if token == nil {
			m.WithErr(global.JwtParseError, c)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if token.Valid {
			if ok {
				c.Set("userUid", claims["Uid"])
				c.Set("userPassword", claims["Password"])
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorMalformed != 0 {
				m.WithErr(global.IncorrectToken, c)
			} else if ve.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
				m.WithErr(global.JwtExpiredError, c)
			} else {
				m.WithErr(global.JwtParseError, c)
			}
			c.Abort()
		}

		c.Next()
	}
}