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

	// trustedAddrs, err := utils.GetWgAddrs()
	// if err != nil {
	// 	log.Printf("Fail to get address of %s: %v", utils.IntfName, err)
	// }

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.SetTrustedProxies(trustedAddrs)
	r.POST("/keepalive", handlers.KeepAliveHandler)

	r.Run(fmt.Sprintf("0.0.0.0:%d", device.ListenPort))
}
