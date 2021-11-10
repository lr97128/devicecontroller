package main

import (
	"flag"
	"fmt"
	"log"
)

var firstKey = "0123456789asdfgh"
var secret = "0123456789asdfgh"

func main() {
	var path, plainPass string
	var encrypt bool
	flag.StringVar(&path, "config", "./config.json", "Where is the configuration file")
	flag.BoolVar(&encrypt, "en", false, "Do you want to encrypt to plain password")
	flag.StringVar(&plainPass, "passwd", "", "Input plain password to encrypt")
	flag.Parse()
	if encrypt {
		if plainPass != "" {
			ciphPass, err := GetCiphPass(plainPass, firstKey, secret)
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("encrypt password is:", ciphPass)
		} else {
			log.Fatal("Please input your plain password")
			return
		}
	} else {
		devs, err := GetDevices(path)
		if err != nil {
			log.Fatal("Get devices from config file failed. err:", err)
		}
		for _, device := range devs {
			if err := device.InitDevice(firstKey); err != nil {
				log.Fatalf("decrypte CinperPassword failed,err is:%v\n", err)
			}
			defer device.Close()
			ret, err := device.Execute()
			if err != nil {
				log.Fatalf("execute command failed,err is:%v\n", err)
			}
			fmt.Printf("%v\n", ret)
		}
	}
}
