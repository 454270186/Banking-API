package domain

import (
	"RESTful/logger"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type AuthRepo interface {
	IsAuthorized(token string, vars gin.Params) bool
}

type RemoteAuthRepo struct{}

func NewAuthRepo() AuthRepo {
	return RemoteAuthRepo{}
}

func (r RemoteAuthRepo) IsAuthorized(token string, vars gin.Params) bool {
	// convert gin.Params -> map[string]string
	paramMap := make(map[string]string)
	for _, param := range vars {
		paramMap[param.Key] = param.Value
	}

	u := buildVerifyURL(token, paramMap)
	log.Println(u)
	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		m := map[string]string{}
		// log.Println(response.Body)
		if err := json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server: " + err.Error())
			return false
		}
		log.Println(m["error"])
		return m["Verified"] == "Authorized"
	}
}

func buildVerifyURL(token string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8181", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	for k, v := range vars {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String()
}