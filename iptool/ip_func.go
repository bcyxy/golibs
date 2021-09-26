package iptool

import (
	"fmt"
	"strconv"
	"strings"
)

// IPStr2Int IP字符串转整型
func IPStr2Int(ipStr string) (ipInt uint32, retErr error) {
	ipParts := strings.Split(ipStr, ".")
	if len(ipParts) != 4 {
		retErr = fmt.Errorf("format_err")
		return
	}
	for _, ipPart := range ipParts {
		iv, err := strconv.Atoi(ipPart)
		if err != nil || iv < 0 || iv > 255 {
			retErr = fmt.Errorf("format_err")
			return
		}
		ipInt = ipInt<<8 + uint32(iv)
	}
	return
}

// IPv4Int2Str IP整型转字符串
func IPv4Int2Str(ipInt uint32) (ipStr string) {
	return fmt.Sprintf("%d.%d.%d.%d", ipInt>>24&0xff, ipInt>>16&0xff,
		ipInt>>8&0xff, ipInt&0xff)
}
