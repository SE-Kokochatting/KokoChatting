package utilstruct

import (
	"KokoChatting/model/dataobject"
	"github.com/golang-jwt/jwt/v4"
)

// Claims 用于jwt验证信息
type Claims struct {
	dataobject.UserProfile
	jwt.RegisteredClaims
}