package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)


var key  = []byte("secret")

var (
	TokenExpired     error  = errors.New("token is expired")
	TokenNotValidYet error  = errors.New("token not active yet")
	TokenMalformed   error  = errors.New("that's not even a token")
	TokenInvalid     error  = errors.New("couldn't handle this token")
)

type JWTClaims struct {
	ID 		interface{} 		`bson:"_id"`
	Email 	string		`bson:"email"`
	//Phone 	string		`bson:"phone"`
	jwt.StandardClaims
}

// JWT中间件
func JWTAuth() gin.HandlerFunc  {
	return func(c *gin.Context) {

		clientToken := c.Request.Header.Get("Authorization")
		log.Print(clientToken)
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "jwt must be provided",
				"code": 401,
			})
			c.Abort()
			return
		}
		clientToken = strings.Replace(string(clientToken), "Bearer ", "", 1)
		claims, err := ParseToken(clientToken)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": "token expired",
					"code": 401,
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": err.Error(),
			})
			c.Abort()
			return
		}

		// 继续交由下一个路由处理，并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// 创建token
func CreateToken(claims JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

// 解析token
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token can not be handled")
}

// 更新token
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return CreateToken(*claims)
	}
	return "", TokenInvalid
}
