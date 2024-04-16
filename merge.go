package main

import (
	"cmp"
	"net"
	"slices"
)

func merge(cidrs []net.IPNet) []net.IPNet {
	merged := []net.IPNet{}

	slices.SortFunc[[]net.IPNet, net.IPNet](cidrs, func(a net.IPNet, b net.IPNet) int {
		onesA, _ := a.Mask.Size()
		onesB, _ := b.Mask.Size()
		return cmp.Compare[int](onesA, onesB)
	})

	currentOnes := 0
	moreSpecifics := []net.IPNet{}
	for _, cidr := range cidrs {
		ones, _ := cidr.Mask.Size()
		if ones > currentOnes {
			moreSpecifics = merged
			currentOnes = ones
		}

		if hasMoreSpecific(cidr, moreSpecifics) {
			continue
		}

		merged = append(merged, cidr)
	}

	return merged
}

func hasMoreSpecific(cidr net.IPNet, moreSpecifics []net.IPNet) bool {
	for _, ms := range moreSpecifics {
		if ms.Contains(cidr.IP) {
			return true
		}
	}

	return false
}
