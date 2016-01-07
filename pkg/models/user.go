package models

import (
	"code.google.com/p/go-uuid/uuid"
	"errors"
	//"fmt"
)

func UserAdd(u *User) error {
	ut := User{}
	err := DB.One("user", "WHERE name='"+u.Name+"'", &ut, []string{"tel"})
	if err == nil {
		return errors.New("登陆账户为" + u.Name + "的账户已经存在")
	}

	u.Id = uuid.NewUUID().String()
	err = DB.Insert("user", u)
	return err
}
