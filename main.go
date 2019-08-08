package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/hongjie104/NAS-server/app/pkg/config"
	"github.com/hongjie104/NAS-server/app/routers"
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

	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf("%s", config.Config.Server.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
