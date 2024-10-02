package main

import (
	"cmp"
	"net"
	"slices"
)

func Merge(cidrs []net.IPNet) []net.IPNet {
	merged := []net.IPNet{}

	slices.SortFunc[[]net.IPNet, net.IPNet](cidrs, func(a net.IPNet, b net.IPNet) int {
		onesA, _ := a.Mask.Size()
		onesB, _ := b.Mask.Size()
		return cmp.Compare[int](onesA, onesB)
	})

	currentOnes := 0
	lessSpecifics := []net.IPNet{}
	for _, cidr := range cidrs {
		ones, _ := cidr.Mask.Size()
		if ones > currentOnes {
			lessSpecifics = merged
			currentOnes = ones
		}

		if hasLessSpecific(cidr, lessSpecifics) {
			continue
		}

		merged = append(merged, cidr)
	}

	return merged
}

func hasLessSpecific(cidr net.IPNet, lessSpecifics []net.IPNet) bool {
	for _, ms := range lessSpecifics {
		if ms.Contains(cidr.IP) {
			return true
		}
	}

	return false
}
