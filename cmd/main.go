package main

import (
	"fmt"
	"workshop/confReading"
)

func main() {
	prepare()
	// startserv()
}

func prepare() {
	conf := confReading.New()

	conf.LoadConfig()

	fmt.Println(*conf.GetConf())
	//db connection

}
