package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId uint) (string, error){
	// トークンの有効期限を設定
	// 今回は1時間
	expire := time.Now().Add(time.Hour * 1).Unix()
	// トークンのペイロードを設定
	// 今回はユーザーIDと有効期限を設定
	claims := jwt.MapClaims{
		"userId": userId,
		"exp": expire,
	}
	// トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 環境変数からシークレットキーを取得
	secretKey := os.Getenv("SECRET_KEY")
	// トークンを署名
	tokenString, err := token.SignedString([]byte(secretKey)); if err != nil {
		return "", err
	}
	return tokenString, nil
}