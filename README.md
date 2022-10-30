# wg-nat

wg-nat is a tool helping WireGuard peers behind NAT or firewall to communicate as well as turning tunnels into a full-mesh network.

<p align="middle">
<img src="https://raw.githubusercontent.com/Nativu5/nativu5.github.io/main/files/202210291101014.svg" width=45% title="Centralized network" />
&nbsp; => &nbsp;
<img src="https://raw.githubusercontent.com/Nativu5/nativu5.github.io/main/files/202210291101012.svg" width=45% title="Using wg-nat to make full mesh network" />
</p>

# Concept

The tunnel between the registry server and client (e.g., client A) is guaranteed by persistent keepalives. This application multiplexes the port used by that tunnel so that other clients could use it to directly connect to client A.

The registry server is also responsible for informing other clients about client A's endpoint information.

Inspirations are from: 
* [WireGuard Endpoint Discovery and NAT Traversal using DNS-SD](https://www.jordanwhited.com/posts/wireguard-endpoint-discovery-nat-traversal/)

* [wireguard-tools nat-hole-punching example](https://git.zx2c4.com/wireguard-tools/tree/contrib/nat-hole-punching)

# Usage 

To utilize this application, a machine with internet access and public IP address is required for **registry**.

Machines behind the firewall or NAT could be configured as clients.

## Initialize Wireguard interfaces

Set up Wireguard interfaces as usual. A registry interface is set and make sure that all clients have configurations to connect with the registry interface (as shown in the first figure).

## Run the Server and Client application

To download the newest binaries supporting multiple platforms, check [GitHub Action Artifacts](https://github.com/Nativu5/wg-nat/actions).  

The server application collects endpoint information and distributes that to all clients so that they can connect with each other. 

```
Usage of server:
  -i string
        Interface name to use (default "wg0")
```

* `-i` must be followed with the name of Wireguard interface. Note that the interface should be brought up **before** the server application launching.


```
Usage of client:
  -i string
        Interface name to use (default "wg0")
  -r string
        Registry public key
  -t duration
        Time interval to send keepalive (default 1m0s)
```

* `-i` is already mentioned in the server usage.  
* `-r` is used to distinguish the registry server from peers by specifying the server's public key.
* `-t` is the time interval for the client to actively communicate with the server. The value should be given in Golang time format.

Run the server and client application with proper arguments. Then the client should be able to connect to other peers directly.

# Security

Be awared that the server should only be accessed within the your WireGuard network. It is very important to set up firewall rules to avoid exposing server on Internet. 