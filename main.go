package main

import (
	"flag"
	"fmt"

	"github.com/atultherajput/go_crash_course/net"
)

var appType = flag.String("app", "grpc", "Type of application")
var port = flag.Int("port", 8080, "The server port")

func main() {
	flag.Parse()

	switch *appType {
	case "mux":
		fmt.Println("Running Mux app")
		net.RunMux(port)
	case "gin":
		fmt.Println("Running Gin app")
		net.RunGin(port)
	case "grpc":
		fmt.Println("Running GRPC app")
		net.RunGrpc(port)
	case "muxpq":
		fmt.Println("Running MuxPq app")
		net.RunMuxPq(port)
	default:
		fmt.Println("Running GRPC app")
		net.RunGrpc(port)
	}

}
