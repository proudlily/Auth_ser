package models

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func List(table string, req string, data *string, sort string) error {
	rows, err := SQLDB.Query(`SELECT d FROM ` + table + ` ` + req + ` ` + sort)
	if err != nil {
		fmt.Println("该处需要错误日志系统", err)
		return err
	}
	var d string
	*data = "["
	for rows.Next() {
		err = rows.Scan(&d)
		*data = *data + d + `,`
	}
	rows.Close()
	*data = strings.TrimRight(*data, ",")
	*data = *data + "]"
	return nil
}

func ListLmit(table string, req string, data *string, sort string, p int, limit int8) {
	rows, err := SQLDB.Query(`SELECT d FROM `+table+` `+req+` `+sort+` LIMIT $1 OFFSET $2`, limit, int(limit)*(p-1))
	if err != nil {
		fmt.Println("该处需要错误日志系统", err)
		return
	}
	var d string
	*data = "["
	for rows.Next() {
		err = rows.Scan(&d)
		*data = *data + d + `,`
	}
	rows.Close()
	*data = strings.TrimRight(*data, ",")
	*data = *data + "]"
}

func GET(table string, req string, re *string) (err error) {
	err = SQLDB.QueryRow(`SELECT d FROM ` + table + ` ` + req).Scan(re)
	if err != nil {
		return errors.New("查询出错，该错误已记录，请联系管理员")
	}
	return
}

func Insert(table string, data *map[string]interface{}) (err error) {
	uid := uuid.NewUUID()
	str := base64.URLEncoding.EncodeToString([]byte(uid))
	(*data)["id"] = str[:len(str)-2]

	jsonString, _ := json.Marshal(data)
	_, err = SQLDB.Exec(`INSERT INTO `+table+` (d) VALUES($1)`, string(jsonString))
	if err != nil {
		return
	}
	return
}

func InsertAuto(table string, data *map[string]interface{}) (err error) {
	jsonString, _ := json.Marshal(data)
	_, err = SQLDB.Exec(`INSERT INTO `+table+` (d) VALUES($1)`, string(jsonString))
	if err != nil {
		return
	}
	return
}

func UpdateAll(table string, data *map[string]interface{}) (err error) {
	var jsonString []byte
	jsonString, err = json.Marshal(data)
	_, err = SQLDB.Exec(`UPDATE `+table+` SET d =$1 WHERE  d @> '{"id": "`+(*data)["id"].(string)+`"}'`, string(jsonString))
	if err != nil {
		return
	}
	return
}

func Update(table string, req string, data *map[string]interface{}) (err error) {
	var jsonString []byte
	jsonString, err = json.Marshal(data)
	_, err = SQLDB.Exec(`UPDATE `+table+` SET d =jsonb_merge(d,$1) `+req, string(jsonString))
	if err != nil {
		return
	}
	return
}

func Inc(table string, req string, key string, num string) (err error) {
	_, err = SQLDB.Exec(`UPDATE ` + table + ` SET d = jsonb_inc(d,'` + key + `',` + num + `) ` + req)
	if err != nil {
		return
	}
	return
}

func Del(table string, req string) (err error) {
	_, err = SQLDB.Exec(`DELETE FROM ` + table + ` ` + req)
	if err != nil {
		return
	}
	return
}

func Count(table string, req string) int64 {
	var re int64
	err := SQLDB.QueryRow(`SELECT COUNT(*) FROM ` + table + ` ` + req).Scan(&re)
	if err != nil {
		return 0
	}
	return re
}
