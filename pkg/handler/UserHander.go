package handler

import (
	"fmt"
	//"github.com/asyoume/Auth/thrift_go/User"
)

type UserHandler struct {
}

func (this UserHandler) Register(u string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}

func (this UserHandler) Login(key, pwd string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}

func (this UserHandler) Logout(key string) (r string, err error) {
	fmt.Print("ping()\n")
	return "", nil
}
