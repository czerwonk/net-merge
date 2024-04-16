package main

import "net"

func parseCIDR(s string) (valid bool, res *net.IPNet) {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		return parseIP(s)
	}

	return true, ipnet
}

func parseIP(s string) (valid bool, res *net.IPNet) {
	ip := net.ParseIP(s)
	if ip == nil {
		return false, nil
	}

	return true, ipToIPNet(ip)
}

func ipToIPNet(ip net.IP) *net.IPNet {
	if ip.To4() == nil {
		return &net.IPNet{
			IP:   ip,
			Mask: net.CIDRMask(128, 128),
		}
	}

	return &net.IPNet{
		IP:   ip,
		Mask: net.CIDRMask(32, 32),
	}
}
