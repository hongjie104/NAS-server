package main

import (
	"github.com/hongjie104/NAS-server/app/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run("127.0.0.1:7001")
}
