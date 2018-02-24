package main

import (
	"../Service"
	"runtime"
)

func main() {

	Service.RunMyEchoService()
	runtime.Goexit()
}
