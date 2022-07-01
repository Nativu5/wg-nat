package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var (
	client    *wgctrl.Client
	intfName  *string
	regPubkey *string
)

func KeepAlive(registry *net.UDPAddr, intfPubkey string) error {
	url := "http://" + registry.String() + "/keepalive"
	param := "?pubkey=" + intfPubkey

	resp, err := http.Post(url+param, "text/plain", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var newPeerConfig []wgtypes.PeerConfig
	if err := json.Unmarshal(body, &newPeerConfig); err != nil {
		return err
	}

	if err := client.ConfigureDevice(*intfName, wgtypes.Config{
		ReplacePeers: true,
		Peers:        newPeerConfig,
	}); err != nil {
		return err
	}

	return nil
}

func main() {
	intfName = flag.String("interface", "wg0", "Interface name to use")
	regPubkey = flag.String("registry", "", "Registry public key")
	flag.Parse()

	log.Println("Hello, World!")

	var err error
	client, err = wgctrl.New()
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	device, err := client.Device(*intfName)
	if err != nil {
		log.Panic(err)
	}

	var registry wgtypes.Peer
	for _, v := range device.Peers {
		if v.PublicKey.String() == *regPubkey {
			registry = v
			break
		}
	}

	if registry.AllowedIPs == nil {
		log.Panic("No valid peer is found")
	}

	regAddr := &net.UDPAddr{
		IP:   registry.AllowedIPs[0].IP,
		Port: registry.Endpoint.Port,
	}

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		KeepAlive(regAddr, device.PublicKey.String())
	}
}
