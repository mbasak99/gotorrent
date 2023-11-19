package peers

import "net"

type Peer struct {
	IP   net.IP
	Port uint8
}
