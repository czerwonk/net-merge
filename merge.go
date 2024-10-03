package main

import (
	"cmp"
	"net"
	"slices"
	"sync"

	"github.com/infobloxopen/go-trees/iptree"
)

func Merge(cidrs []net.IPNet) []net.IPNet {
	v4 := []net.IPNet{}
	v6 := []net.IPNet{}
	for _, cidr := range cidrs {
		if cidr.IP.To4() == nil {
			v6 = append(v6, cidr)
		} else {
			v4 = append(v4, cidr)
		}
	}

	ch := make(chan net.IPNet)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go mergeSubset(v4, ch, wg)
	go mergeSubset(v6, ch, wg)
	go func() {
		wg.Wait()
		close(ch)
	}()

	merged := []net.IPNet{}
	for cidr := range ch {
		merged = append(merged, cidr)
	}

	return merged
}

func mergeSubset(cidrs []net.IPNet, ch chan<- net.IPNet, wg *sync.WaitGroup) {
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
		ch <- cidr
	}

	wg.Done()
}
