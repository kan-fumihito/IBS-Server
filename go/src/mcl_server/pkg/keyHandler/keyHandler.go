package keyHandler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Firebase SDK のセットアップ

		opt := option.WithCredentialsFile("/go/src/mcl_server/pkg/keyHandler/secretkey.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		log.Print(opt)
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		// クライアントから送られてきた JWT 取得
		authHeader := c.QueryParam("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWT が無効なら Handler に進まず別処理
			fmt.Printf("error verifying ID token: %v\n", err)
			return c.String(http.StatusUnauthorized, "error verifying ID token\n")
		}
		log.Printf("Verified ID token: %v\n", token)
		return next(c)
	}
}

func getParamFromJwt(param string, tokenString string) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})
	if err != nil {
		return ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims[param].(string)
	} else {
		return ""
	}
}

func Secretkey(c echo.Context) error {
	authHeader := c.QueryParam("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	email := getParamFromJwt("email", idToken)

	return c.String(http.StatusOK, getSecretKey(time.Now().Format("2006-01-02 15:04:05"), email))
}

func Secretkey2(c echo.Context) error {
	authHeader := c.QueryParam("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	email := getParamFromJwt("email", idToken)

	return c.String(http.StatusOK, getSecretKey2(time.Now().Format("2006-01-02 15:04:05"), email))
}

func Publickey(c echo.Context) error {
	authHeader := c.QueryParam("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	p1 := getParamFromJwt("p1", idToken)
	t := getParamFromJwt("time", idToken)

	return c.String(http.StatusOK, getPublicKey(t, p1))
}

func Publickey2(c echo.Context) error {
	authHeader := c.QueryParam("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	p2 := getParamFromJwt("p2", idToken)
	t := getParamFromJwt("time", idToken)

	return c.String(http.StatusOK, getPublicKey2(t, p2))
}
