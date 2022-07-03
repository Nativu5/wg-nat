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
		log.Fatal("Failed to initialize WireGuard.")
	}
	defer utils.Client.Close()

	device, err := utils.Client.Device(utils.IntfName)
	if err != nil {
		log.Fatal(err)
	}

	localIP, err := utils.GetWgLocalIP()
	if err != nil {
		log.Fatal(err)
	}

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.POST("/keepalive", handlers.KeepAliveHandler)

	r.Run(fmt.Sprintf("%s:%d", localIP, device.ListenPort))
}
