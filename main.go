package main

import (
	"flag"
	"github.com/Yesterday17/go-drcom-jlu/drcom"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	activeMAC = ""
	client    *drcom.Client
	cfg       *drcom.Config
)

// return code list
// 10 failed to parse config file

func main() {
	var cfgPath string
	var err error

	flag.StringVar(&cfgPath, "config", "./config.json", "配置文件的路径")
	flag.Parse()

	Interfaces = make(map[string]*Interface)

	if err = initWireless(); err != nil {
		log.Fatal(err)
	}

	if err = initWired(); err != nil {
		log.Fatal(err)
	}

	// 加载配置文件
	cfg, err = ReadConfig(cfgPath)
	if err != nil {
		log.Println(err)
		os.Exit(10)
	}

	var MAC string = cfg.MAC

	go watchNetStatus()

	if MAC != "" {
		NewClient(MAC)
	}

	// 处理退出信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sig
		log.Printf("Exiting with signal %s", s.String())
		if activeMAC != "" && client != nil {
			// client.Logout()
			_ = client.Close()
		}
		return
	}
}
