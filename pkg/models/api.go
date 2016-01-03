package models

import (
	"fmt"
	"github.com/asyoume/Auth/pkg/log"
	db_drive "github.com/asyoume/postgres"
)

func Init() {
	//初始化数据库
	sqlurl := "postgres://postgres:all123456@m.aolilong.net/wxmall?sslmode=disable"
	loger := log.KafakLoger{}
	err := db_drive.Init(sqlurl, &loger)
	if err != nil {
		fmt.Println("初始化postgresql数据库失败")
	} else {
		fmt.Println("初始化postgresql数据库成功")
	}

	test := Test1{}
	err = db_drive.One("test1", "", &test, []string{"test1", "test2"})
	fmt.Println(err)
	fmt.Println(test)

	tests := []Test1{}
	err = db_drive.All("test1", "", &tests, []string{"test1", "test2"}, "", 1, 7)
	fmt.Println(err)
	fmt.Println(tests)
}
