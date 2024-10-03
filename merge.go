package main

import (
	"net"
	"slices"
	"sync"

	"github.com/infobloxopen/go-trees/iptree"
	"golang.org/x/exp/maps"
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
	defer wg.Done()

	grouped := groupCIDRsByPrefixLength(cidrs)
	if len(grouped) == 0 {
		return
	}

	keys := maps.Keys(grouped)
	asc := func(a, b int) int {
		return a - b
	}
	slices.SortStableFunc(keys, asc)
	maxPfxLen := keys[len(keys)-1]

	t := iptree.NewTree()

	for _, k := range keys {
		vals := grouped[k]

		for _, cidr := range vals {
			if _, found := t.GetByNet(cidr); found {
				continue
			}

			if k < maxPfxLen {
				t = t.InsertNet(cidr, nil)
			}
			ch <- *cidr
		}
	}
}

func groupCIDRsByPrefixLength(cidrs []net.IPNet) map[int][]*net.IPNet {
	grouped := make(map[int][]*net.IPNet)

	for _, cidr := range cidrs {
		pfxL, _ := cidr.Mask.Size()
		grouped[pfxL] = append(grouped[pfxL], &cidr)
	}

	return grouped
}
