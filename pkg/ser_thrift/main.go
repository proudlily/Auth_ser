package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/asyoume/Auth_ser/pkg/handler"
	"os"
	"path/filepath"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	transport := flag.String("transport", "binary", "Specify transport (framed, buffered, file, memory, zlib)")

	buffered := flag.String("buffered", "off", "Use buffered transport")
	framed := flag.String("framed", "off", "Use framed transport")

	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	secure := flag.Bool("secure", false, "Use tls secure transport")

	conf_path := flag.String("conf", "nil", "配置文件路径")

	flag.Parse()

	if *conf_path == "nil" {
		Usage()
		os.Exit(1)
	}

	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory

	if *buffered == "on" {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if *framed == "on" {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	//获取配置文件的路径
	complete_path(conf_path)
	//初始化控制层
	handler.Init(*conf_path)

	if err := runServer(transportFactory, *transport, protocolFactory, *addr, *secure); err != nil {
		fmt.Println("error running server:", err)
	}

}

func complete_path(conf_path *string) {
	if len((*conf_path)) > 0 && string((*conf_path)[0]) != "/" {
		var curr_path string = ""
		curr_path, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		*conf_path = curr_path + "/" + *conf_path
	}
}
