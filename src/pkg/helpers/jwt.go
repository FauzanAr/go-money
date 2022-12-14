package helper

import (
	"money-management/src/config"
	"money-management/src/pkg/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtHelper interface {
	Generate(u string, e string, uI string) utils.Result
	Validate(token string) utils.Result
	Refresh() utils.Result
}

type JwtClaims struct {
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	UserId		string	`json:"id"`
	jwt.StandardClaims
}

type JwtConfig struct {
	secretKey	string
	issuer		string
}

type JwtResult struct {
	AccessToken		string	`json:"access_token"`
	RefreshToken	string	`json:"refresh_token"`
}

func Jwt() JwtHelper {
	return &JwtConfig{
		secretKey: config.Get().SecretKey,
		issuer: config.Get().Issuer,
	}
}

func (j *JwtConfig) Generate(username string, email string, userId string) utils.Result {
	var res utils.Result
	tokenClaims := JwtClaims{
		Username: username,
		Email: email,
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Issuer: j.issuer,
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}

	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenClaims)
	token, err := plainToken.SignedString([]byte(j.secretKey))
	if err != nil {
		Logger.Error("Error while creating JWT token, msg: " + err.Error())
		res.Error = err
		return res
	}

	refreshTokenClaims := jwt.StandardClaims{
		Issuer: j.issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		IssuedAt: time.Now().Unix(),
	}

	plainRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, refreshTokenClaims)
	refreshToken, err := plainRefreshToken.SignedString([]byte(j.secretKey))
	if err != nil {
		Logger.Error("Error while creating JWT refresh token, msg: " + err.Error())
		res.Error = err
		return res
	}

	result := JwtResult{
		AccessToken: token,
		RefreshToken: refreshToken,
	}
	res.Data = result
	return res
}

func (j *JwtConfig) Validate(token string) utils.Result {
	var res utils.Result
	var claims JwtClaims
	parseToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	switch errType := err.(type) {
	case jwt.ValidationError:
		if errType.Errors == jwt.ValidationErrorExpired {
			res.Error = "token expired"
			return res
		} else {
			res.Error = "token parse failed"
		}
	}

	if !parseToken.Valid {
		res.Error = "token parse failed"
		return res
	}

	res.Data = claims
	return res
}

func (j *JwtConfig) Refresh() utils.Result {
	var res utils.Result

	return res
}