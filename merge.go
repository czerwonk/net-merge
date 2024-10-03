package main

import (
	"cmp"
	"net"
	"slices"

	"github.com/infobloxopen/go-trees/iptree"
)

func Merge(cidrs []net.IPNet) []net.IPNet {
	merged := []net.IPNet{}

	slices.SortFunc[[]net.IPNet, net.IPNet](cidrs, func(a net.IPNet, b net.IPNet) int {
		onesA, _ := a.Mask.Size()
		onesB, _ := b.Mask.Size()
		return cmp.Compare[int](onesA, onesB)
	})

	t := iptree.NewTree()
	for _, cidr := range cidrs {
		if _, found := t.GetByNet(&cidr); found {
			continue
		}

		t = t.InsertNet(&cidr, nil)
		merged = append(merged, cidr)
	}

	return merged
}
