package utils

import (
	"flag"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var (
	IntfName string
	Client   *wgctrl.Client
	Device   *wgtypes.Device
)

func init() {
	flag.StringVar(&IntfName, "interface", "wg0", "Interface name to use")
	flag.Parse()
}

func InitWg() error {
	var err error
	Client, err = wgctrl.New()
	if err != nil {
		return err
	}

	Device, err = Client.Device(IntfName)
	if err != nil {
		return err
	}
	return nil
}

func GeneratePeerConfig(peer wgtypes.Peer) wgtypes.PeerConfig {
	return wgtypes.PeerConfig{
		PublicKey:                   peer.PublicKey,
		Remove:                      false,
		PresharedKey:                &peer.PresharedKey, // ?
		UpdateOnly:                  true,
		Endpoint:                    peer.Endpoint,
		ReplaceAllowedIPs:           false,
		AllowedIPs:                  peer.AllowedIPs,
		PersistentKeepaliveInterval: nil,
	}
}
