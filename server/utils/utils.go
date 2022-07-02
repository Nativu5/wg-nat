package utils

import (
	"flag"
	"net"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var (
	IntfName string
	Client   *wgctrl.Client
)

func init() {
	flag.StringVar(&IntfName, "i", "wg0", "Interface name to use")
	flag.Parse()
}

func InitWg() error {
	var err error
	Client, err = wgctrl.New()
	return err
}

func GeneratePeerConfig(peer wgtypes.Peer) wgtypes.PeerConfig {
	return wgtypes.PeerConfig{
		PublicKey:                   peer.PublicKey,
		Remove:                      false,
		PresharedKey:                &peer.PresharedKey,
		UpdateOnly:                  false,
		Endpoint:                    peer.Endpoint,
		ReplaceAllowedIPs:           false,
		AllowedIPs:                  peer.AllowedIPs,
		PersistentKeepaliveInterval: &peer.PersistentKeepaliveInterval,
	}
}

func GetWgAddrs() ([]string, error) {
	wgIntf, err := net.InterfaceByName(IntfName)
	if err != nil {
		return nil, err
	}

	wgAddrs, _ := wgIntf.Addrs()
	addrs := make([]string, 0, len(wgAddrs))
	for _, v := range wgAddrs {
		addrs = append(addrs, v.String())
	}
	return addrs, nil
}

func SaveConfigToFile() error {
	return nil
}
