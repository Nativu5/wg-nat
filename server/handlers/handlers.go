package handlers

import (
	"log"
	"net/http"

	"github.com/Nativu5/wg-nat/server/utils"
	"github.com/gin-gonic/gin"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// KeepAliveHandler handles the keepalive request
// and reply with known peers' configuration.
func KeepAliveHandler(ctx *gin.Context) {
	pubkey := ctx.Query("pubkey")

	device, err := utils.Client.Device(utils.IntfName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Println(err)
		return
	}

	var sender wgtypes.Peer
	newPeerConfig := make([]wgtypes.PeerConfig, 0, len(device.Peers))
	for _, v := range device.Peers {
		if v.PublicKey.String() == pubkey {
			// skip keepalive sender's config
			sender = v
			continue
		}
		if v.Endpoint == nil {
			// skip unknown endpoint
			continue
		}
		newPeerConfig = append(newPeerConfig, utils.GeneratePeerConfig(v))
	}

	if sender.Endpoint == nil {
		// no matching peer found
		ctx.JSON(http.StatusBadRequest, nil)
		log.Printf("No matching peer found for %s", pubkey)
		return
	}

	log.Printf("New endpoint of peer %s is %s", pubkey, sender.Endpoint.String())
	ctx.JSON(http.StatusOK, newPeerConfig)
}
