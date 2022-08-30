package handler

import (
	"database/sql"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type NewHandler struct {
	DB *sql.DB
}

var mySigningKey = []byte("unicorns")

func (h NewHandler)GetJWT(key string) (string, error) {
	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Get JWT ", "time":"%s"}`, time.Now()))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Krissanawat"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = key + "secret"

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
