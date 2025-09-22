package main

import (
	"log"
	"net"

	"github.com/google/nftables"
	"github.com/google/nftables/expr"
	"golang.org/x/sys/unix"
)

func main() {
	c := &nftables.Conn{}

	// สร้าง table inet filter (ถ้ายังไม่มี)
	table := &nftables.Table{
		Family: nftables.TableFamilyINet,
		Name:   "filter",
	}
	c.AddTable(table)

	// สร้าง chain input (hook: input)
	chain := &nftables.Chain{
		Name:     "input",
		Table:    table,
		Type:     nftables.ChainTypeFilter,
		Hooknum:  nftables.ChainHookInput,
		Priority: nftables.ChainPriorityFilter,
	}
	c.AddChain(chain)

	// ตัวอย่าง : อนุญาต TCP port 443 จาก 192.168.1.0/24
	_, ipNet, _ := net.ParseCIDR("192.168.1.0/24")
	role := &nftables.Rule{
		Table: table,
		Chain: chain,
		Exprs: []expr.Any{
			// payload load inet_proto
			&expr.Meta{
				Key:      expr.MetaKeyL4PROTO,
				Register: 1,
			},
			// cmp eq tcp (6)
			&expr.Cmp{
				Op:       expr.CmpOpEq,
				Register: 1,
				Data:     []byte{unix.IPPROTO_TCP},
			},
			// payload load destination port
			&expr.Payload{
				DestRegister: 1,
				Base:         expr.PayloadBaseTransportHeader,
				Offset:       2, // offset for TCP dest port
				Len:          2,
			},
			// cmp dst port == 443
			&expr.Cmp{
				Op:       expr.CmpOpEq,
				Register: 1,
				Data:     []byte{0x01, 0xbb}, // 443 in hex
			},

			// match source IP
			&expr.Payload{
				DestRegister: 2,
				Base:         expr.PayloadBaseNetworkHeader,
				Offset:       12, // IPV4 src offset
				Len:          4,
			},
			&expr.Cmp{
				Op:       expr.CmpOpEq,
				Register: 2,
				Data:     ipNet.IP.To4(),
			},
			// verdict accept
			&expr.Verdict{
				Kind: expr.VerdictAccept,
			},
		},
	}
	c.AddRule(role)
	if err := c.Flush(); err != nil {
		log.Fatalf("nft flush: %v", err)
	}
	log.Println("nft rule added (allow TCP 443 from 192.168.1.0/24)")
}
