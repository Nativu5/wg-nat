package handlers

import (
	"net/http"

	"github.com/Nativu5/wg-nat/server/utils"
	"github.com/gin-gonic/gin"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func KeepAliveHandler(ctx *gin.Context) {
	pubkey := ctx.Query("pubkey")
	if pubkey == "" {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	// 似无意义
	var peer wgtypes.Peer
	for _, v := range utils.Device.Peers {
		if v.PublicKey.String() == pubkey {
			peer = v
			break
		}
	}
	if peer.Endpoint == nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	newPeerConfig := make([]wgtypes.PeerConfig, 0, len(utils.Device.Peers))
	for _, v := range utils.Device.Peers {
		newPeerConfig = append(newPeerConfig, utils.GeneratePeerConfig(v))
	}

	ctx.JSON(http.StatusOK, newPeerConfig)
}
