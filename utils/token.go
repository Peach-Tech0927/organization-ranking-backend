package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenFromId(id uint) (string, error) {
	tokenLifeSpanStr := os.Getenv("TOKEN_LIFE_SPAN")
	if len(tokenLifeSpanStr) == 0 {
		return "", errors.New("TOKEN_LIFE_SPAN is not set in the environment")
	}
	tokenLifeSpan, err := strconv.Atoi(tokenLifeSpanStr)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"id": id,
		"exp": time.Now().Add(time.Hour * time.Duration(tokenLifeSpan)).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func TokenValid(c *gin.Context) error {
	tokenStr, err := getTokenStringFromRequestHeader(c)
	if err != nil {
		return err
	}

	token, err := parseToken(tokenStr)

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token is not valid")
	}

	return nil
}

func ExtractTokenId(c *gin.Context) (uint, error) {
	tokenStr, err := getTokenStringFromRequestHeader(c)
	if err != nil {
		return 0, err
	}

	token, err := parseToken(tokenStr)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("error while parsing claims")
	}

	id, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("error while parsing id")
	}

	return uint(id), nil
}

func getTokenStringFromRequestHeader(c *gin.Context) (string, error) {
	bearToken := c.Request.Header.Get("Authorization")
    strArr := strings.Split(bearToken, " ")
    if len(strArr) == 2 {
        return strArr[1], nil
    }

    return "", errors.New("no token found")
}

func parseToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("there was an error while parsing the token")
        }
        return []byte(os.Getenv("API_SECRET")), nil
    })

    if err != nil {
        return nil, err
    }

    return token, nil
}