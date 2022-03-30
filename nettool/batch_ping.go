package nettool

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type PingItem struct {
	IP     string
	SendTs int64 // 发包时间
	RcvTs  int64 // 收包时间
}

// BatchPing 批量ping
func BatchPing(ctx context.Context, ips []string) (<-chan PingItem, error) {
	retCh := make(chan PingItem, 10)
	ipMap := make(map[string]*PingItem)
	for _, ip := range ips {
		ipMap[ip] = &PingItem{IP: ip}
	}

	// 创建连接，此处需要root用户
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return retCh, fmt.Errorf("listen_failed:%v", err)
	}

	// go 收包，收包中判断结束
	go func() {
		defer close(retCh)

		msg := []byte{}
		cnt := 0
		for {
			_, sIP, err := conn.ReadFrom(msg)
			if err != nil { // conn被关闭
				return
			}

			itPtr, ok := ipMap[sIP.String()]
			if !ok || itPtr.RcvTs != 0 { // 非本次发出的IP or 重复返回
				continue
			}
			itPtr.RcvTs = time.Now().UnixMilli()
			retCh <- *itPtr
			if cnt++; cnt >= len(ipMap) {
				return
			}
		}
	}()

	// go 发包
	go func() {
		defer conn.Close()

		pid := os.Getpid()
		seq := 1
		tkr := time.NewTicker(time.Second)
		defer tkr.Stop()
		for {
			for ip, itPtr := range ipMap {
				if itPtr.RcvTs != 0 {
					continue
				}

				dst, err := net.ResolveIPAddr("ip", ip)
				msg := &icmp.Message{
					Type: ipv4.ICMPTypeEcho,
					Code: 0,
					Body: &icmp.Echo{
						ID:   pid,
						Seq:  seq,
						Data: []byte("abcde"),
					},
				}
				seq++
				msgBytes, _ := msg.Marshal(nil)
				_, err = conn.WriteTo(msgBytes, dst)
				if err != nil { // conn被关闭
					return
				}
				if itPtr.SendTs == 0 {
					itPtr.SendTs = time.Now().UnixMilli()
				}
			}
			select {
			case <-ctx.Done():
				return
			case <-tkr.C:
			}
		}
	}()

	return retCh, nil
}
