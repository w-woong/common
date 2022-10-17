package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenClaim struct {
	UsrID  string  `json:"usr_id"`
	AccLvl string  `json:"admin"`
	Sub    float64 `json:"sub"`
	Exp    float64 `json:"exp"`
}

// GenerateAccessToken 액세스토큰을 생성
func GenerateAccessToken(tokenSecret string, usrID string, accLvl string, sub float64, expMinute time.Duration) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["usr_id"] = usrID
	claims["acc_lvl"] = accLvl
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(expMinute).Unix()

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// GenerateRefreshToken 액세스토큰을 재생성할 수 있는 리프레시토큰 생성
func GenerateRefreshToken(tokenSecret, usrID string, sub float64, expMinute time.Duration) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["usr_id"] = usrID
	rtClaims["sub"] = sub
	rtClaims["exp"] = time.Now().Add(expMinute).Unix()

	rt, err := refreshToken.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return rt, nil
}

// ValidateAccessToken 액세스토큰의 유효성을 판단하고 리프레시 가능한지 확인
func ValidateAccessToken(accessToken string, tokenSecret string) (*TokenClaim, bool, error) {
	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(tokenSecret), nil
	})

	var claim *TokenClaim
	if token != nil && token.Claims.Valid() != nil {
		c := token.Claims.(jwt.MapClaims)
		claim = &TokenClaim{
			UsrID:  c["usr_id"].(string),
			AccLvl: c["acc_lvl"].(string),
			Sub:    c["sub"].(float64),
			Exp:    c["exp"].(float64),
		}
	}
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			return claim, true, err
		}
		return nil, false, err
	}

	if token.Valid {
		c := token.Claims.(jwt.MapClaims)
		claim = &TokenClaim{
			UsrID:  c["usr_id"].(string),
			AccLvl: c["acc_lvl"].(string),
			Sub:    c["sub"].(float64),
			Exp:    c["exp"].(float64),
		}
		return claim, false, nil
	}

	return nil, false, errors.New("not authorized")
}

// ValidateRefreshToken 리프레시토큰의 유효성 검증
func ValidateRefreshToken(refreshToken, tokenSecret string, usrID string, sub float64) (float64, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(tokenSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		// if claims["sub"].(float64) == sub {
		// 	return nil
		// }
		rtUsrID, ok := claims["usr_id"].(string)
		if !ok {
			return 0, errors.New("not authorized")
		}
		rtSub, ok := claims["sub"].(float64)
		if !ok {
			return 0, errors.New("not authorized")
		}
		if rtUsrID == usrID {
			return rtSub, nil
		}

		return 0, errors.New("not authorized")
	}

	return 0, err
}
