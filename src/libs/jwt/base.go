package jwt

import (
	"be-blog/src/constants"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func CreateExpTime(exp string) time.Time {
	switch exp {
	case constants.ONE_HOUR:
		return time.Now().Add(time.Hour)
	case constants.SEVEN_DAY:
		return time.Now().Add(time.Hour * 24 * 7)
	case constants.SIX_MONTH:
		return time.Now().Add(time.Hour * 24 * 30 * 6)
	case constants.ONE_YEAR:
		return time.Now().Add(time.Hour * 24 * 365)
	case constants.TEN_MIN:
		return time.Now().Add(time.Minute * 10)
	default:
		return time.Now().Add(time.Hour)
	}
}
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	secret := viper.GetString("jwt.secret")
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token không hợp lệ")
}

func CreateTokenUser(payload PayloadUser, exp string) (string, error) {
	secret := viper.GetString("jwt.secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        payload.ID,
		"full_name": payload.FullName,
		"user_name": payload.UserName,
		"email":     payload.Email,
		"avatar":    payload.Avatar,
		"roles":     payload.Roles,
		"exp":       CreateExpTime(exp).Unix(),
	})
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ParseTokenPayloadUser(tokenString string) (*PayloadUser, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return &PayloadUser{
		ID:       claims["id"].(string),
		FullName: claims["full_name"].(string),
		UserName: claims["user_name"].(string),
		Email:    claims["email"].(string),
		Roles:    claims["roles"].(string),
		Avatar:   claims["avatar"].(string),
		Exp:      int64(claims["exp"].(float64)),
	}, nil
}

func CreateTokenVerifyEmail(payload PayloadVerify, exp string) (string, error) {
	secret := viper.GetString("jwt.secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": payload.Email,
		"exp":   CreateExpTime(exp).Unix(),
	})
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ParseTokenPayloadVerifyEmail(tokenString string) (*PayloadVerify, error) {
	clain, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return &PayloadVerify{
		Email: clain["email"].(string),
		Exp:   int64(clain["exp"].(float64)),
	}, nil
}

func CreateForgotPassToken(payload PayloadResetPass, exp string) (string, error) {
	secret := viper.GetString("jwt.secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": payload.Email,
		"exp":   CreateExpTime(exp).Unix(),
	})
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ParseTokenPayloadForgotPass(tokenString string) (*PayloadResetPass, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return &PayloadResetPass{
		Email: claims["email"].(string),
		Exp:   int64(claims["exp"].(float64)),
	}, nil
}
