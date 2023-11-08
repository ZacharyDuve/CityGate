package host

import (
	"net"
	"time"
)

type Host interface {
	Addr() net.Addr
	SubDomain() string
	Expiration() time.Time
}
