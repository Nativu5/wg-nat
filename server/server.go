package main

import (
	"fmt"
	"log"

	"github.com/Nativu5/wg-nat/server/handlers"
	"github.com/Nativu5/wg-nat/server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	if utils.InitWg() != nil {
		log.Fatal("Failed to initialize WireGuard")
	}

	r := gin.Default()
	r.POST("/keepalive", handlers.KeepAliveHandler)

	r.Run(fmt.Sprintf("0.0.0.0:%d", utils.Device.ListenPort))
}
