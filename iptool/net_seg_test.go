package iptool_test

import (
	"fmt"
	"testing"

	"github.com/bcyxy/golibs/iptool"
)

func TestNetSegs(t *testing.T) {
	nsMap := iptool.NewNetSegMap()
	nsMap.AddNetSeg("192.168.11.0/24")
	nsMap.AddNetSeg("192.168.0.0/16")
	nsMap.AddNetSeg("10.122.20.0/24")

	netSegs, _ := nsMap.GetNetSegs("192.168.11.2")
	fmt.Println(netSegs)
}
