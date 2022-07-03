package utils

import (
	"flag"
	"fmt"
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

// GeneratePeerConfig generates a PeerConfig from a Peer.
func GeneratePeerConfig(peer wgtypes.Peer) wgtypes.PeerConfig {
	// Client may add new peers or update existed ones.
	// However, removal of peers is not supported.
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

// GetWgCIDR returns the list of WireGuard interface's CIDR.
func GetWgCIDR() ([]string, error) {
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

// GetWgLocalIP returns the local IP of WireGuard interface.
func GetWgLocalIP() (string, error) {
	wgIntf, err := net.InterfaceByName(IntfName)
	if err != nil {
		return "", err
	}

	wgAddrs, _ := wgIntf.Addrs()
	for _, v := range wgAddrs {
		if v.(*net.IPNet).IP != nil {
			return v.(*net.IPNet).IP.String(), nil
		}
	}
	return "", fmt.Errorf("no valid Local IP found")
}
