package utils

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"common-web-framework/response"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const encrypt = "this is encrypt"

// CreateAccessToken 创建Token
func CreateAccessToken(id int, username string) response.TokenResponse {

	var create = time.Now()

	var expire = time.Now().Add(common.TokenExpire)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expire.Unix(),
		IssuedAt:  create.Unix(),
		Id:        username,
		Issuer:    username,
		Subject:   strconv.Itoa(id),
	})

	accessToken, err := claims.SignedString([]byte(encrypt))

	if err != nil {
		helper.ErrorToResponse(common.CreateTokenFail)
	}

	return response.TokenResponse{
		Token:  accessToken,
		Expire: FormatDate(expire),
		Create: FormatDate(create),
	}

}

// ParseTokenToUserId 解析Token
func ParseTokenToUserId(token string) int {

	claims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(encrypt), nil
	})

	if err != nil || !claims.Valid {
		return -1
	}

	var uid, _ = strconv.Atoi(claims.Claims.(*jwt.StandardClaims).Subject)

	return uid
}
