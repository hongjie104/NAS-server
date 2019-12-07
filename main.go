package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/hongjie104/NAS-server/config"
	"github.com/hongjie104/NAS-server/pkg/log"
	"github.com/hongjie104/NAS-server/router"
)

func main() {
	// r := routers.InitRouter()
	// r.Run(":7002")
	// s := &http.Server{
	// 	Addr:           setting.HTTPPort,
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

	// /*
	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf("%s", config.Config.Server.HTTPPort)

	server := endless.NewServer(endPoint, router.InitRouter())
	server.BeforeBegin = func(add string) {
		log.LogInfo(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}

	server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGTERM] = append(
		server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGTERM],
		preSIGTERM)
	server.SignalHooks[endless.POST_SIGNAL][syscall.SIGTERM] = append(
		server.SignalHooks[endless.POST_SIGNAL][syscall.SIGTERM],
		postSIGTERM)

	err := server.ListenAndServe()
	if err != nil {
		log.LogError(fmt.Sprintf("Server err: %v", err))
	}
	//*/
}

func preSIGTERM() {
	fmt.Println("pre SIGTERM")
}

func postSIGTERM() {
	fmt.Println("post SIGTERM")
}
