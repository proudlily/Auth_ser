package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
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

	flag.Parse()

	var protocolFactory thrift.TProtocolFactory
	fmt.Println("protocol:", *protocol)
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
		fmt.Println("buffered:", *buffered)
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if *framed == "on" {
		fmt.Println("framed:", *framed)
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	if err := runServer(transportFactory, *transport, protocolFactory, *addr, *secure); err != nil {
		fmt.Println("error running server:", err)
	}

}
