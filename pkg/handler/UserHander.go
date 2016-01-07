package handler

import (
	"encoding/json"
	"fmt"
	"github.com/asyoume/Auth/pkg/models"
)

type UserHandler struct {
}

func (this UserHandler) Register(u, device_id string) (r string, err error) {
	user := models.User{}
	err = json.Unmarshal([]byte(u), &user)
	if err != nil {
		return r, err
	}
	err = models.UserAdd(&user)
	if err != nil {
		return r, err
	}
	return user.Id, nil
}

func (this UserHandler) Login(key, pwd, device_id string) (r string, err error) {
	err = models.UserCheck(&user)
	if err != nil {
		return r, err
	}
	return user.Id, nil
	return "", nil
}

func (this UserHandler) Logout(key, device_id string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}

func (this UserHandler) AppRegister(u, device_id, token string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}

func (this UserHandler) AppLogin(key, pwd, device_id, token string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}

func (this UserHandler) AccessToken(appid, secret string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}

func (this UserHandler) Authorize(appid, response_type, scope, state string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}

func (this UserHandler) Oauth2Token(appid, secret, code string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}

func (this UserHandler) UserInfo(token, appid, openid string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}
