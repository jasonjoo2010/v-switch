package main

import (
	co "V-switch/conf"
	aes "V-switch/crypt"
	pl "V-switch/plane"
	tap "V-switch/tap"
	to "V-switch/tools"
	_u "V-switch/udp"
	"fmt"
	"log"
	"os"
)

func init() {

	if (os.Getuid() != 0) && (os.Getgid() != 0) {
		fmt.Println("[OMG] YOU MUST BE ROOT TO RUN THIS PROGRAM , BECAUSE ONLY ROOT IS ALLOWED TO CREATE TAP DEVICES")
		os.Exit(1)
	}

	to.LogEngineStart()
	co.StartConfig()
	tap.EngineStart()
	pl.PlaneInit()
	aes.EngineStart()
	_u.UDPEngineStart()

}

func main() {

	log.Println("[OMG] End of bootstrap.")

	select {}

	os.Exit(0)

}
