package main

import (
	"fmt"
	_ "github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/jackpal/gateway"
	"net"
)

type str1 struct {
	UserId     int64 `json:"user_id"`
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type"`
}

type str2 struct {
	UserId   int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
}

func main() {
	gatewayIP, err := gateway.DiscoverGateway()
	if err != nil {
		fmt.Println("Error finding gateway:", err)
		return
	}

	fmt.Println("Gateway / Router IP Address:", gatewayIP)

	localIPs, err := localIPs(gatewayIP)
	if err != nil {
		fmt.Println("Error finding local IP addresses:", err)
		return
	}

	for _, ip := range localIPs {
		fmt.Println("Local IP Address:", ip)
	}
}

func localIPs(gatewayIP net.IP) ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var ips []net.IP
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		ip := ipNet.IP.To4()
		if ip == nil || ip.IsLoopback() || ip.Equal(gatewayIP) {
			continue
		}
		ips = append(ips, ip)
	}

	return ips, nil
}
