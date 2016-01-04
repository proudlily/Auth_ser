package main

import (
	"flag"
	"fmt"
	"github.com/asyoume/Auth/pkg/handler"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"os"
	"path/filepath"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

var uhander = handler.UserHandler{}

func main() {
	//获取执行参数
	conf_path := flag.String("conf", "", "配置文件路径")
	flag.Parse()
	//获取配置文件的路径
	if len((*conf_path)) > 0 && string((*conf_path)[0]) != "/" {
		var curr_path string = ""
		curr_path, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		*conf_path = curr_path + "/" + *conf_path
	}
	//初始化控制层
	handler.Init(*conf_path)

	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Use(cors.Default().Handler)

	// Routes
	e.Post("/Login", UserLogin)
	e.Get("/Register", UserRegister)
	e.Run(":9091")
}
