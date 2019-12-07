package nft

import (
	"github.com/google/nftables"
)

var conn *nftables.Conn

func Connect() *nftables.Conn {
	conn = &nftables.Conn{}
	return conn
}
