package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var path string
	flag.StringVar(&path, "config", "./config.json", "Where is the configuration file")
	flag.Parse()
	devs, err := GetDevices(path)
	if err != nil {
		log.Fatal("Get devices from config file failed. err:", err)
	}
	for _, device := range devs {
		plainPass, err := GetPlainPass(device.CiphPass, device.Secret)
		if err != nil {
			log.Fatal("Descrypt CiphPass failed, err:", err)
			return
		}
		device.PlainPass = plainPass
		fmt.Println(device)
	}
}
