package utilstruct

import (
	"github.com/golang-jwt/jwt/v4"
)

// Claims 用于jwt验证信息
type Claims struct {
	Uid uint64
	Password string
	jwt.RegisteredClaims
}