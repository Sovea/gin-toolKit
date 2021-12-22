package service_jwt

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetUserId needs a *gin.Content and it will returns current user's ID.
func GetUserId(c *gin.Context) (string, error) {
	//Version 1
	// token := c.Request.Header.Get("Authorization")
	// claims, err := jwt_ex.ParseToken(token)
	// if err !=nil{
	// 	return "",err
	// }else{
	// 	return claims.ID,nil
	// }

	//Version 2
	claims := c.MustGet("claims").(*CustomClaims)
	return claims.ID, nil

}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": -1,
				"msg":    "no access",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		j := NewJWT()
		// parseToken
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": -1,
					"msg":    "authorization has expired",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

// JWT struct
type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "75sa16ve036my654wor84ld"
)

type CustomClaims struct {
	ID       string `json:"userId"`
	Name     string `json:"name"`
	RealName string `json:"realname"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return SignKey
}

func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// parse Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
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
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// update token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func (j *JWT) RefreshTokenWithOther(tokenString string, new_columns map[string]string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		mutable := reflect.ValueOf(claims).Elem()
		for k, v := range new_columns {
			mutable.FieldByName(k).SetString(v)
		}
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
