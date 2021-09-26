package iptool

import (
	"fmt"
	"net"
)

// NetSegLib 网段映射表
type NetSegLib struct {
	segCache map[int]*map[uint32]bool
}

// NewNetSegMap 创建
func NewNetSegMap() *NetSegLib {
	return &NetSegLib{
		segCache: make(map[int]*map[uint32]bool),
	}
}

// AddNetSeg 添加网段
func (sf *NetSegLib) AddNetSeg(netSeg string) error {
	_, ipNet, err := net.ParseCIDR(netSeg)
	if err != nil {
		return err
	}

	ones, _ := ipNet.Mask.Size()
	netPtr, ok := sf.segCache[ones]
	if !ok {
		netPtr = &map[uint32]bool{}
		sf.segCache[ones] = netPtr
	}

	ipInt, err := IPStr2Int(ipNet.IP.String())
	if err != nil {
		return err
	}
	(*netPtr)[ipInt] = true

	return nil
}

// GetNetSegs 获取IP归属的网段
func (sf *NetSegLib) GetNetSegs(ip string) (netSegs []string, retErr error) {
	ipInt, err := IPStr2Int(ip)
	if err != nil {
		retErr = err
		return
	}
	for mask, netMap := range sf.segCache {
		yw := 32 - mask
		net := ipInt >> yw << yw
		_, ok := (*netMap)[net]
		if ok {
			netSegs = append(netSegs, fmt.Sprintf("%v/%v", IPv4Int2Str(net), mask))
		}
	}
	return
}
