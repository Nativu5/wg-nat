# wg-nat

wg-nat is a tool helping WireGuard peers behind NAT or firewall to communicate as well as turning tunnels into a full-mesh network.

# Usage 

To utilize this application, a machine with internet access and public IP address is required. A **registry interface** and **server application** are configured on it.

Machines behind the firewall or NAT could be configured as clients. A Wireguard interface directly connected with the registry is set and the client application runs on it. 


## Initialize Wireguard interfaces

Set up Wireguard interfaces as usual. A registry interface is required, which should contain desired peer configurations. And make sure that the client has a proper configuration to connect with the registry interface.

## Run the Server and Client application

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