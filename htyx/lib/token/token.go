package token

import (
	"errors"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ErrOther         = errors.New("Other Error")
	ErrMissingHeader = errors.New("The length of the `Token` header is zero.")
)

type PayLoad struct {
	Uid uint64
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*PayLoad, error) {
	pd := &PayLoad{}
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return pd, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		pd.Uid = uint64(claims["uid"].(float64))
		return pd, nil
	} else {
		return pd, ErrOther
	}
}
func ParseRequest(c *gin.Context) (*PayLoad, error) {
	token := c.GetHeader("token")
	secret := viper.GetString("jwt_secret")
	if len(token) == 0 {
		return &PayLoad{}, ErrMissingHeader
	}
	return Parse(token, secret)

}

func Sign(pl PayLoad, secret string) (tokenstring string, err error) {
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": pl.Uid,
		"nbf": time.Now().Unix(),

		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenstring, err = token.SignedString([]byte(secret))
	return

}
